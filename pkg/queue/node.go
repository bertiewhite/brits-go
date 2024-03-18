package queue

type node[T any] struct {
	Data T
	Next *node[T]
	Prev *node[T]
}

// this is a bad name but I like it so it stays!!
func nodify[T any](data T) node[T] {
	return node[T]{
		Data: data,
	}
}
