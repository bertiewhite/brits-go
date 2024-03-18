package queue

// this could be some fun stuff or haeve some channels for sending/ consuming events associated with the queue but it's all a bunch of fun

type Window interface {
	// add more methods to let you see
	QueueLength() int
}

type window struct {
	queueLength *int
}

func (w *window) QueueLength() int { return *w.queueLength }
