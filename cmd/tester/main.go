package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/bertiewhite/brits-go/pkg/testharness"
)

func main() {

	testHarness, err := testharness.New("localhost:8000")
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	tmpCount := -1
	sentCount := &tmpCount

	go func() {
		defer wg.Done()
		count, err := testHarness.StartReceiving(context.Background(), sentCount)
		if err != nil {
			fmt.Println("Failure to recvieve message", err)
			return
		}
		fmt.Printf("Received %d messages\n", count)
	}()

	count, err := testHarness.StartSending(context.Background(), 2*time.Minute)
	if err != nil {
		fmt.Println("Failure to send message", err)
		return
	}
	fmt.Printf("Sent     %d messages\n", count)

	*sentCount = count

	wg.Wait()
}
