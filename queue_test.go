package queue

import (
	"errors"
	"sync"
	"testing"
)

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
		v, err := q.Peek()
		if errors.Is(err, ErrEmpty) {
			break
		}
		if err != nil {
			t.Fatalf("expect nil err, got %v", err)
		}
		if _, ok := testData[v]; !ok {
			t.Fatalf("unexpect %d return", v)
		}
		v, err = q.Pop()
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
	_, err := q.Pop()
	if err == nil {
		t.Fatal("expect err, got nil")
	}
	if !errors.Is(err, ErrEmpty) {
		t.Fatalf("expect %v, got %v", ErrEmpty, err)
	}
}
