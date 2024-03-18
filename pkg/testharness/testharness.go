package testharness

import (
	"context"
	"fmt"
	"time"

	"github.com/bertiewhite/brits-go/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TestHarness struct {
	client proto.MessageQueueClient
}

var logAfter = 50000

func New(addr string) (*TestHarness, error) {

	// security schmeurity
	grpcOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(addr, grpcOpts...)
	if err != nil {
		return nil, err
	}

	return &TestHarness{
		client: proto.NewMessageQueueClient(conn),
	}, nil
}

func (t *TestHarness) StartSending(ctx context.Context, duration time.Duration) (int, error) {
	sender, err := t.client.Send(ctx)
	if err != nil {
		return 0, err
	}
	defer func() {
		err := sender.CloseSend()
		if err != nil {
			fmt.Printf("error closing sender: %v\n", err)
		}
	}()

	timeout := time.After(duration)
	count := 0

	for {
		select {
		case <-timeout:
			return count, nil
		default:
			if count%logAfter == 0 {
				fmt.Printf("Have sent %d messages\n", count)
			}
			err := sender.Send(&proto.MessagePayload{Data: []byte("hello")})
			if err != nil {
				return 0, err
			}
			count++
		}
	}
}

func (t *TestHarness) StartReceiving(ctx context.Context, sentCount *int) (int, error) {
	if sentCount == nil {
		return 0, fmt.Errorf("sentCount must be non-nil")
	}

	receiver, err := t.client.Receive(ctx, &proto.Empty{})
	if err != nil {
		return 0, err
	}
	defer func() {
		err := receiver.CloseSend()
		if err != nil {
			fmt.Printf("error closing receiver: %v\n", err)
		}
	}()

	rcvCount := 0
	notLogged := true

	for {
		if rcvCount%logAfter == 0 {
			fmt.Printf("Have received %d messages\n", rcvCount)
		}
		_, err := receiver.Recv()
		if err != nil {
			return 0, err
		}
		rcvCount++

		if *sentCount > 0 && notLogged {
			fmt.Println("Sent count is ", *sentCount)
			notLogged = false
		}
		if *sentCount > 0 && rcvCount >= *sentCount {
			return rcvCount, nil
		}
	}
}
