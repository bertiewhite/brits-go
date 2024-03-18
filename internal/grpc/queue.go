package grpc

import (
	"errors"
	"time"

	grpc "github.com/bertiewhite/brits-go/pkg/proto"
	"github.com/bertiewhite/brits-go/pkg/queue"
)

type QueueService struct {
	grpc.UnimplementedMessageQueueServer

	queue *queue.Queue[[]byte]
}

func NewQueueService() *QueueService {
	return &QueueService{queue: queue.NewQueue[[]byte]()}
}

func (s *QueueService) Send(stream grpc.MessageQueue_SendServer) error {
	for {
		msg, err := stream.Recv()
		if err != nil {
			// I should update this to return some useful info. But not now
			stream.SendAndClose(&grpc.Empty{})
			return err
		}
		s.queue.Add(msg.Data)
	}
}

func (s *QueueService) Receive(_ *grpc.Empty, stream grpc.MessageQueue_ReceiveServer) error {
	var outErr error
	// this is for now it can be quicker than ~60 msg a minute
LOOP:
	for {
		data, err := s.queue.Take()
		switch {
		case errors.Is(err, queue.ErrQueueIsEmpty):
			time.Sleep(1 * time.Second)
			continue
		case err != nil:
			// this is unnecessary but I like it to make it explicit
			outErr = err
			break LOOP
		}
		err = stream.Send(&grpc.MessagePayload{Data: data})
		if err != nil {
			return err
		}
	}

	return outErr
}
