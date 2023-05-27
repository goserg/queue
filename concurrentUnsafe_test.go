package queue

import (
	"testing"
)

func TestZeroConcurrentUnsafeQueue_Push(t *testing.T) {
	o := NewConcurrentUnsafe[int](0)
	o.Push(777)
	popWithCheck(t, o, 777)
}

func TestConcurrentUnsafeQueue_Push(t *testing.T) {
	testData := []int{12, 34, 5, 2, 5, 65, 123, 65, 3, 4, 32, 54, 12, 76, 78, 3, 2, 15}
	q := NewConcurrentUnsafe[int](3)
	for _, val := range testData {
		q.Push(val)
		popWithCheck(t, q, val)
	}
}

func TestConcurrentUnsafeQueue_ScaleUp(t *testing.T) {
	testData := []int{12, 34, 5, 2, 5, 65, 123, 65, 3, 4, 32, 54, 12, 76, 78, 3, 2, 15}
	q := NewConcurrentUnsafe[int](3)
	for _, val := range testData {
		q.Push(val)
	}
	for _, val := range testData {
		popWithCheck(t, q, val)
	}
}

func TestConcurrentUnsafeQueue_ScaleUpWithRotation(t *testing.T) {
	testBatch1 := []int{12, 34, 685}
	testBatch2 := []int{5, 2, 5}
	testBatch3 := []int{5, 65, 123, 63, 56}
	q := NewConcurrentUnsafe[int](7)
	for _, i := range testBatch1 {
		q.Push(i)
	}
	for _, i := range testBatch2 {
		q.Push(i)
	}
	for _, i := range testBatch1 {
		popWithCheck(t, q, i)
	}
	for _, i := range testBatch3 {
		q.Push(i)
	}
	for _, i := range testBatch2 {
		popWithCheck(t, q, i)
	}
	for _, i := range testBatch3 {
		popWithCheck(t, q, i)
	}
}

func TestConcurrentUnsafeQueue_Peek(t *testing.T) {
	testBatch1 := []int{12, 34, 685}
	testBatch2 := []int{5, 2, 5}
	testBatch3 := []int{5, 65, 123, 63, 56}
	q := NewConcurrentUnsafe[int](7)
	for i, val := range testBatch1 {
		l := q.Len()
		if l != i {
			t.Fatalf("expexted len %d, got %d", i, l)
		}
		q.Push(val)
	}
	for i, val := range testBatch2 {
		l := q.Len()
		if l != i+len(testBatch1) {
			t.Fatalf("expexted len %d, got %d", i, l)
		}
		q.Push(val)
	}
	for i := range testBatch1 {
		l := q.Len()
		if l != len(testBatch1)+len(testBatch2)-i {
			t.Fatalf("expexted len %d, got %d", i, l)
		}
		q.Pop()
	}
	for i, val := range testBatch3 {
		l := q.Len()
		if l != len(testBatch2)+i {
			t.Fatalf("expexted len %d, got %d", i, l)
		}
		q.Push(val)
	}
}

func TestConcurrentUnsafeQueue_Len(t *testing.T) {
	testData := []int{12, 34, 5, 2, 5, 65, 123, 65, 3, 4, 32, 54, 12, 76, 78, 3, 2, 15}
	q := NewConcurrentUnsafe[int](3)
	for _, val := range testData {
		q.Push(val)
	}
}

func popWithCheck[T comparable](t *testing.T, q *ConcurrentUnsafeQueue[T], expected T) {
	t.Helper()
	v, err := q.Pop()
	if err != nil {
		t.Fatalf("expect nil err, got %v", err)
	}
	if v != expected {
		t.Fatalf("expect %v got %v", expected, v)
	}
}

func peekWithCheck[T comparable](t *testing.T, q *ConcurrentUnsafeQueue[T], expected T) {
	t.Helper()
	v, err := q.Peek()
	if err != nil {
		t.Fatalf("expect nil err, got %v", err)
	}
	if v != expected {
		t.Fatalf("expect %v got %v", expected, v)
	}
}
