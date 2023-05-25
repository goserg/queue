package queue

import (
	"errors"
	"sync"
)

type Queue[T any] struct {
	mu      sync.Mutex
	storage []T
	next    int
	tail    int
}

var ErrEmpty = errors.New("empty queue")

func New[T any](size int) *Queue[T] {
	return &Queue[T]{
		storage: make([]T, size),
	}
}

func (q *Queue[T]) Push(value T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if (q.next+1)%len(q.storage) == q.tail {
		newStore := make([]T, len(q.storage)*2)
		for i := 0; i < len(q.storage); i++ {
			v, err := q.pop()
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

func (q *Queue[T]) Pop() (T, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.pop()
}

func (q *Queue[T]) pop() (T, error) {
	var zero T
	if q.next == q.tail {
		return zero, ErrEmpty
	}
	val := q.storage[q.tail]
	q.tail = (q.tail + 1) % len(q.storage)
	return val, nil
}
