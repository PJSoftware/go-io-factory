package factory

type Queue[T any] struct {
	pointersToT []*T
}

// NewQueue[T]() constructs a new queue holding (pointers to) elements of type T
func NewQueue[T any]() *Queue[T] {
	q := &Queue[T]{
		pointersToT: []*T{},
	}

	return q
}

// q.Push(T) pushes a new element of type T to the tail of the queue. It
// actually just stores a pointer to the element, since that is what we need to
// return later.
func (q *Queue[T]) Push(element T) {
	q.pointersToT = append(q.pointersToT, &element)
}

// q.Next() removes the head element of the queue, and returns a pointer to it.
// If the queue is empty it returns nil.
func (q *Queue[T]) Next() *T {
	if q.IsEmpty() {
		return nil
	}
	elPtr := q.pointersToT[0]
	q.pointersToT = q.pointersToT[1:]
	return elPtr
}

// q.Peek() returns a pointer to the head element of the queue, without removing
// it from the queue. If the queue is empty it returns nil.
func (q Queue[T]) Peek() *T {
	if q.IsEmpty() {
		return nil
	}
	return q.pointersToT[0]
}

func (q Queue[T]) IsEmpty() bool {
	return len(q.pointersToT) == 0
}
