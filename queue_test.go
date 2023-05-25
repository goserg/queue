package queue

import (
	"errors"
	"sync"
	"testing"
)

func TestZeroQueue_Push(t *testing.T) {
	o := New[int](0)
	o.Push(777)
	v, err := o.Pop()
	if err != nil {
		t.Fatalf("expect nil err, got %v", err)
	}
	if v != 777 {
		t.Fatalf("expect %d got %d", 777, v)
	}
}

func TestQueue_Push(t *testing.T) {
	testData := []int{12, 34, 5, 2, 5, 65, 123, 65, 3, 4, 32, 54, 12, 76, 78, 3, 2, 15}
	q := New[int](3)
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

func TestQueue_ScaleUp(t *testing.T) {
	testData := []int{12, 34, 5, 2, 5, 65, 123, 65, 3, 4, 32, 54, 12, 76, 78, 3, 2, 15}
	q := New[int](3)
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

func TestConcurrent(t *testing.T) {
	testData := map[int]struct{}{
		2:   {},
		96:  {},
		31:  {},
		464: {},
		64:  {},
		4:   {},
		86:  {},
		35:  {},
		5:   {},
		87:  {},
		826: {},
	}
	q := New[int](3)
	var wg sync.WaitGroup
	for val := range testData {
		val := val
		wg.Add(1)
		go func() {
			q.Push(val)
			wg.Done()
		}()
	}
	wg.Wait()
	for {
		v, err := q.Pop()
		if errors.Is(err, ErrEmpty) {
			break
		}
		if err != nil {
			t.Fatalf("expect nil err, got %v", err)
		}
		if _, ok := testData[v]; !ok {
			t.Fatalf("unexpect %d return", v)
		}
		delete(testData, v)
	}
	if len(testData) != 0 {
		t.Fatalf("missing elemets: %v", testData)
	}
}
