package factory

type Queue[T any] struct {
	contents []T
	emptyVal T
}

// NewQueue[T]() constructs a new queue holding elements of type T. It allows us
// to specify the default value to be returned if the queue is empty.
func NewQueue[T any](empty T) *Queue[T] {
	q := &Queue[T]{
		contents: []T{},
		emptyVal: empty,
	}

	return q
}

// q.Push(T) pushes a new element of type T to the tail of the queue.
func (q *Queue[T]) Push(element T) {
	q.contents = append(q.contents, element)
}

// q.Next() removes the head element of the queue, and returns it. If the queue is empty it returns the predefined default value.
func (q *Queue[T]) Next() T {
	if q.IsEmpty() {
		return q.emptyVal
	}
	elPtr := q.contents[0]
	q.contents = q.contents[1:]
	return elPtr
}

// q.Peek() returns the head element of the queue, without removing it from the
// queue. If the queue is empty it returns the predefined default value.
func (q Queue[T]) Peek() T {
	if q.IsEmpty() {
		return q.emptyVal
	}
	return q.contents[0]
}

// q.IsEmpty() returns true if the queue is empty
func (q Queue[T]) IsEmpty() bool {
	return len(q.contents) == 0
}

// q.PeekAll returns a copy of the entire contents of the queue as a slice
// without actually modifying the slice
func (q Queue[T]) PeekAll() []T {
	all := []T{}
	return append(all, q.contents...)
}

// q.Flush() removes the contents of the queue
func (q *Queue[T]) Flush() {
	q.contents = []T{}
}
