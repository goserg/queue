package queue

import (
	"errors"
	"sync"
)

type Queue[T any] struct {
	mu     sync.Mutex
	unsafe *ConcurrentUnsafeQueue[T]
}

var ErrEmpty = errors.New("empty queue")

func New[T any](size int) *Queue[T] {
	return &Queue[T]{
		unsafe: NewConcurrentUnsafe[T](size),
	}
}

func (q *Queue[T]) Push(value T) {
	q.mu.Lock()
	q.unsafe.Push(value)
	q.mu.Unlock()
}

func (q *Queue[T]) Pop() (T, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.unsafe.Pop()
}

func (q *Queue[T]) Peek() (T, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.unsafe.Peek()
}
