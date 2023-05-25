package queue

import (
	"testing"
)

func TestConcurrentUnsafeQueue_Push(t *testing.T) {
	testData := []int{12, 34, 5, 2, 5, 65, 123, 65, 3, 4, 32, 54, 12, 76, 78, 3, 2, 15}
	q := NewConcurrentUnsafe[int](3)
	for _, val := range testData {
		q.Push(val)
		v, err := q.Pop()
		if err != nil {
			t.Fatalf("expect nil err, got %v", err)
		}
		if v != val {
			t.Fatalf("expect %d got %d", val, v)
		}
	}
}

func TestConcurrentUnsafeQueue_ScaleUp(t *testing.T) {
	testData := []int{12, 34, 5, 2, 5, 65, 123, 65, 3, 4, 32, 54, 12, 76, 78, 3, 2, 15}
	q := NewConcurrentUnsafe[int](3)
	for _, val := range testData {
		q.Push(val)
	}
	for _, val := range testData {
		v, err := q.Pop()
		if err != nil {
			t.Fatalf("expect nil err, got %v", err)
		}
		if v != val {
			t.Fatalf("expect %d got %d", val, v)
		}
	}
}
