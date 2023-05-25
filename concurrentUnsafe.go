package queue

import (
	"errors"
)

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
		for i := 0; i < len(q.storage); i++ {
			v, err := q.Pop()
			if err != nil {
				if errors.Is(err, ErrEmpty) {
					break
				}
				panic(err) // Should not happen
			}
			newStore[i] = v
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
