package queue

type ConcurrentUnsafeQueue[T any] struct {
	storage []T
	next    int
	tail    int
}

func NewConcurrentUnsafe[T any](size int) *ConcurrentUnsafeQueue[T] {
	if size < 1 {
		size = 1
	}
	return &ConcurrentUnsafeQueue[T]{
		storage: make([]T, size),
	}
}

func (q *ConcurrentUnsafeQueue[T]) Push(value T) {
	if (q.next+1)%len(q.storage) == q.tail {
		newStore := make([]T, len(q.storage)*2)
		copy(newStore, q.storage[q.tail:])
		if q.tail > q.next {
			copy(newStore[len(q.storage)-q.tail:], q.storage[:q.next])
		}
		q.next = len(q.storage) - 1
		q.storage = newStore
		q.tail = 0
	}
	q.storage[q.next] = value
	q.next++
	q.next = q.next % len(q.storage)
}

func (q *ConcurrentUnsafeQueue[T]) Pop() (T, error) {
	var zero T
	if q.next == q.tail {
		return zero, ErrEmpty
	}
	val := q.storage[q.tail]
	q.tail = (q.tail + 1) % len(q.storage)
	return val, nil
}

func (q *ConcurrentUnsafeQueue[T]) Peek() (T, error) {
	var zero T
	if q.next == q.tail {
		return zero, ErrEmpty
	}
	return q.storage[q.tail], nil
}
