package XLL

import (
	"errors"
	"math/rand/v2"
	"sync"
	"testing"
)

func TestInsert(t *testing.T) {
	list := New[int]()

	// Test InsertFront
	if err := list.InsertFront(1); err != nil {
		t.Errorf("InsertFront failed: %v", err)
	}
	if list.Size() != 1 {
		t.Errorf("Expected size 1, got %d", list.Size())
	}

	if err := list.InsertFront(2); err != nil {
		t.Errorf("InsertFront failed: %v", err)
	}
	if list.Size() != 2 {
		t.Errorf("Expected size 2, got %d", list.Size())
	}

	// Test InsertBack
	if err := list.InsertBack(3); err != nil {
		t.Errorf("InsertBack failed: %v", err)
	}
	if list.Size() != 3 {
		t.Errorf("Expected size 3, got %d", list.Size())
	}

	// Verify order
	expected := []int{2, 1, 3}
	i := 0
	err := list.TraverseForward(func(data int) {
		if data != expected[i] {
			t.Errorf("Expected %d at position %d, got %d", expected[i], i, data)
		}
		i++
	})
	if err != nil {
		t.Errorf("TraverseForward failed: %v", err)
	}
}

func TestDelete(t *testing.T) {
	list := New[int]()
	// Insert elements and handle potential errors
	for _, v := range []int{1, 2, 3} {
		if err := list.InsertBack(v); err != nil {
			t.Fatalf("InsertBack failed: %v", err)
		}
	}

	// Test DeleteFront
	if err := list.DeleteFront(); err != nil {
		t.Errorf("DeleteFront failed: %v", err)
	}
	if list.Size() != 2 {
		t.Errorf("Expected size 2 after DeleteFront, got %d", list.Size())
	}

	// Test DeleteBack
	if err := list.DeleteBack(); err != nil {
		t.Errorf("DeleteBack failed: %v", err)
	}
	if list.Size() != 1 {
		t.Errorf("Expected size 1 after DeleteBack, got %d", list.Size())
	}

	// Verify remaining element
	err := list.TraverseForward(func(data int) {
		if data != 2 {
			t.Errorf("Expected remaining element to be 2, got %d", data)
		}
	})
	if err != nil {
		t.Errorf("TraverseForward failed: %v", err)
	}

	// Test deleting last element
	if err := list.DeleteFront(); err != nil {
		t.Errorf("DeleteFront failed: %v", err)
	}
	if list.Size() != 0 {
		t.Errorf("Expected size 0 after deleting last element, got %d", list.Size())
	}

	// Test deleting from empty list
	if err := list.DeleteFront(); !errors.Is(err, ErrEmptyList) {
		t.Errorf("Expected ErrEmptyList, got %v", err)
	}
	if err := list.DeleteBack(); !errors.Is(err, ErrEmptyList) {
		t.Errorf("Expected ErrEmptyList, got %v", err)
	}
}

func TestTraverse(t *testing.T) {
	list := New[int]()
	elements := []int{1, 2, 3, 4, 5}
	for _, e := range elements {
		if err := list.InsertBack(e); err != nil {
			t.Errorf("InsertBack failed: %v", err)
		}
	}

	// Test TraverseForward
	i := 0
	err := list.TraverseForward(func(data int) {
		if data != elements[i] {
			t.Errorf("TraverseForward: Expected %d at position %d, got %d", elements[i], i, data)
		}
		i++
	})
	if err != nil {
		t.Errorf("TraverseForward failed: %v", err)
	}

	// Test TraverseBackward
	i = len(elements) - 1
	err = list.TraverseBackward(func(data int) {
		if data != elements[i] {
			t.Errorf("TraverseBackward: Expected %d at position %d, got %d", elements[i], i, data)
		}
		i--
	})
	if err != nil {
		t.Errorf("TraverseBackward failed: %v", err)
	}
}

func TestEdgeCases(t *testing.T) {
	list := New[int]()

	// Test operations on empty list
	if err := list.DeleteFront(); !errors.Is(err, ErrEmptyList) {
		t.Errorf("Expected ErrEmptyList, got %v", err)
	}
	if err := list.DeleteBack(); !errors.Is(err, ErrEmptyList) {
		t.Errorf("Expected ErrEmptyList, got %v", err)
	}
	if list.Size() != 0 {
		t.Errorf("Expected size 0 for empty list, got %d", list.Size())
	}

	// Test single element
	if err := list.InsertFront(1); err != nil {
		t.Errorf("InsertFront failed: %v", err)
	}
	if list.Size() != 1 {
		t.Errorf("Expected size 1 after inserting single element, got %d", list.Size())
	}

	if err := list.DeleteBack(); err != nil {
		t.Errorf("DeleteBack failed: %v", err)
	}
	if list.Size() != 0 {
		t.Errorf("Expected size 0 after deleting single element, got %d", list.Size())
	}

	// Test inserting after freeing
	if err := list.Free(); err != nil {
		t.Errorf("Free failed: %v", err)
	}
	if err := list.InsertFront(1); !errors.Is(err, ErrFreedList) {
		t.Errorf("Expected ErrFreedList, got %v", err)
	}
}

func TestConcurrentOperations(t *testing.T) {
	list := New[int]()
	n := 1000
	var wg sync.WaitGroup

	// Concurrent insertions
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			var err error
			if val%2 == 0 {
				err = list.InsertFront(val)
			} else {
				err = list.InsertBack(val)
			}
			if err != nil {
				t.Errorf("Concurrent insert failed: %v", err)
			}
		}(i)
	}
	wg.Wait()

	if list.Size() != n {
		t.Errorf("Expected size %d after concurrent insertions, got %d", n, list.Size())
	}

	// Concurrent deletions
	for i := 0; i < n/2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var err error
			if rand.IntN(2) == 0 {
				err = list.DeleteFront()
			} else {
				err = list.DeleteBack()
			}
			if err != nil && !errors.Is(err, ErrEmptyList) {
				t.Errorf("Concurrent delete failed: %v", err)
			}
		}()
	}
	wg.Wait()

	if list.Size() != n/2 {
		t.Errorf("Expected size %d after concurrent deletions, got %d", n/2, list.Size())
	}
}
