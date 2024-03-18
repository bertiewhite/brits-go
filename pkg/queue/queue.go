package queue

type Queue[T any] struct {
	Head *node[T]
	Tail *node[T]

	alert []chan struct{}

	length int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) BlockTillNotEmpty() {
	if !q.Empty() {
		return
	}

	c := make(chan struct{}, 1)
	q.alert = append(q.alert, c)
	<-c
}

func (q *Queue[T]) alertNotEmpty() {
	for _, c := range q.alert {
		select {
		case c <- struct{}{}:
		default:
		}
	}
}

func (q *Queue[T]) Empty() bool {
	return q.Head == nil
}

func (q *Queue[T]) Add(data T) {
	node := nodify(data)
	if q.Empty() {
		nodePtr := &node
		q.Head = nodePtr
		q.Tail = nodePtr
		go func() {
			q.alertNotEmpty()
		}()
		return
	}

	q.Head.Prev = &node
	node.Next = q.Head
	q.Head = &node
	q.length++
}

func (q *Queue[T]) Take() (T, error) {
	var empty T
	if q.Empty() {
		return empty, ErrQueueIsEmpty
	}

	node := *q.Tail

	q.Tail = node.Prev
	if q.Tail == nil {
		q.Head = nil
	}
	q.length--

	return node.Data, nil
}

// func (q *Queue[T]) Length() int {
// 	return q.length
// }

func (q *Queue[T]) GetWindow() Window {
	// I wonder if it's possible for this pointer to change
	return &window{queueLength: &q.length}
}
