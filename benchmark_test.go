package XLL

import (
	"container/list"
	"runtime"
	"testing"
)

// BENCHMARK XLL
func BenchmarkInsertFront(b *testing.B) {
	blist := New[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = blist.InsertFront(i)
	}
}

func BenchmarkInsertBack(b *testing.B) {
	blist := New[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = blist.InsertBack(i)
	}
}

func BenchmarkDeleteFront(b *testing.B) {
	blist := New[int]()
	for i := 0; i < b.N; i++ {
		_ = blist.InsertBack(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = blist.DeleteFront()
	}
}

func BenchmarkDeleteBack(b *testing.B) {
	blist := New[int]()
	for i := 0; i < b.N; i++ {
		_ = blist.InsertBack(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = blist.DeleteBack()
	}
}

func BenchmarkTraverseForward(b *testing.B) {
	blist := New[int]()
	for i := 0; i < 1000; i++ {
		_ = blist.InsertBack(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = blist.TraverseForward(func(data int) {})
	}
}

func BenchmarkTraverseBackward(b *testing.B) {
	blist := New[int]()
	for i := 0; i < 1000; i++ {
		_ = blist.InsertBack(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = blist.TraverseBackward(func(data int) {})
	}
}

func BenchmarkInsertFrontWithOptions(b *testing.B) {
	blist := New[int](
		WithBlockSize[int](2048),
		WithGrowthRate[int](1.5),
		WithInitialCapacity[int](1024),
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = blist.InsertFront(i)
	}
}

func BenchmarkInsertBackWithOptions(b *testing.B) {
	blist := New[int](
		WithBlockSize[int](2048),
		WithGrowthRate[int](1.5),
		WithInitialCapacity[int](1024),
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = blist.InsertBack(i)
	}
}

// Benchmark SLICE
type Slice struct {
	data []int
}

func (s *Slice) Insert(val int) {
	s.data = append(s.data, val)
}

func (s *Slice) Traverse() {
	for range s.data {
		// Do nothing, just traverse
	}
}

// Standard doubly linked list for comparison
type DoublyLinkedList struct {
	list *list.List
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{list: list.New()}
}

func (dll *DoublyLinkedList) Insert(val int) {
	dll.list.PushBack(val)
}

func (dll *DoublyLinkedList) Traverse() {
	for e := dll.list.Front(); e != nil; e = e.Next() {
		// Do nothing, just traverse
	}
}

// Benchmark functions
func BenchmarkSliceInsert(b *testing.B) {
	s := &Slice{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Insert(i)
	}
}

func BenchmarkDoublyLinkedListInsert(b *testing.B) {
	dll := NewDoublyLinkedList()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dll.Insert(i)
	}
}

func BenchmarkXLLInsert(b *testing.B) {
	blist := New[int]()
	defer blist.Free()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = blist.InsertFront(i)
	}
}

func BenchmarkSliceTraverse(b *testing.B) {
	s := &Slice{}
	for i := 0; i < 1000000; i++ {
		s.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Traverse()
	}
}

func BenchmarkDoublyLinkedListTraverse(b *testing.B) {
	dll := NewDoublyLinkedList()
	for i := 0; i < 1000000; i++ {
		dll.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dll.Traverse()
	}
}

func BenchmarkXLLTraverse(b *testing.B) {
	blist := New[int]()
	defer blist.Free()
	for i := 0; i < 1000000; i++ {
		_ = blist.InsertFront(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = blist.TraverseForward(func(int) {})
	}
}

// Memory usage benchmark
func BenchmarkMemoryUsage(b *testing.B) {
	b.Run("Slice", func(b *testing.B) {
		s := &Slice{}
		benchmarkMemory(b, func() {
			for i := 0; i < 1000000; i++ {
				s.Insert(i)
			}
		})
	})

	b.Run("DoublyLinkedList", func(b *testing.B) {
		dll := NewDoublyLinkedList()
		benchmarkMemory(b, func() {
			for i := 0; i < 1000000; i++ {
				dll.Insert(i)
			}
		})
	})

	b.Run("XLL", func(b *testing.B) {
		blist := New[int]()
		defer blist.Free()
		benchmarkMemory(b, func() {
			for i := 0; i < 1000000; i++ {
				_ = blist.InsertFront(i)
			}
		})
	})
}

func benchmarkMemory(b *testing.B, f func()) {
	b.ReportAllocs()
	// Warmup
	f()
	runtime.GC()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	beforeAlloc := m.TotalAlloc
	f()
	runtime.ReadMemStats(&m)
	b.ReportMetric(float64(m.TotalAlloc-beforeAlloc), "B/op")
}

// Additional benchmarks for XLL operations
func BenchmarkXLLInsertBack(b *testing.B) {
	blist := New[int]()
	defer blist.Free()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = blist.InsertBack(i)
	}
}

func BenchmarkXLLDeleteFront(b *testing.B) {
	blist := New[int]()
	defer blist.Free()
	for i := 0; i < b.N; i++ {
		_ = blist.InsertBack(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = blist.DeleteFront()
	}
}

func BenchmarkXLLDeleteBack(b *testing.B) {
	blist := New[int]()
	defer blist.Free()
	for i := 0; i < b.N; i++ {
		_ = blist.InsertBack(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = blist.DeleteBack()
	}
}

func BenchmarkXLLTraverseBackward(b *testing.B) {
	blist := New[int]()
	defer blist.Free()
	for i := 0; i < 1000000; i++ {
		_ = blist.InsertBack(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = blist.TraverseBackward(func(int) {})
	}
}
