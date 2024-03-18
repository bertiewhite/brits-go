package queue

import (
	"errors"
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("Adding an element returns it on a take", func(t *testing.T) {
		q := NewQueue[int]()

		in := 3
		q.Add(in)
		out, err := q.Take()

		if err != nil {
			t.Fatalf("Expected nil err got: %s", err.Error())
		}
		if in != out {
			t.Fatalf("Failed to retrieve an item from the queue. expected: %d, got: %d", in, out)
		}
	})

	t.Run("Empty queue returns empty queue err on take", func(t *testing.T) {
		q := NewQueue[string]()
		_, err := q.Take()
		if err == nil {
			t.Fatal("Expected the Empty Queue error got nil")
		}
		if !errors.Is(err, ErrQueueIsEmpty) {
			t.Fatalf("Expected empty queue error got: %s", err.Error())
		}
	})

	t.Run("Multiple adds are returned in FIFO", func(t *testing.T) {
		q := NewQueue[int]()

		q.Add(1)
		q.Add(2)
		q.Add(3)

		first, err := q.Take()
		if err != nil {
			t.Fatalf("Expected nil err got on first take: %s", err.Error())
		}
		second, err := q.Take()
		if err != nil {
			t.Fatalf("Expected nil err got on second take: %s", err.Error())
		}
		third, err := q.Take()
		if err != nil {
			t.Fatalf("Expected nil err got on third take: %s", err.Error())
		}

		if first != 1 {
			t.Fatalf("Expected first out to be 1 got: %d", first)
		}
		if second != 2 {
			t.Fatalf("Expected second out to be 2 got: %d", second)
		}
		if third != 3 {
			t.Fatalf("Expected third out to be 3 got: %d", third)
		}

		fourth, err := q.Take()
		if fourth != 0 {
			t.Fatalf("Expected fourth out to be 0 got: %d", fourth)
		}
		if err == nil {
			t.Fatalf("Expected empty queue error got nil")
		}
		if !errors.Is(err, ErrQueueIsEmpty) {
			t.Fatalf("Expected empty queue error got: %s", err.Error())
		}
	})
}
