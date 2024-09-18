package XLL

import (
	"errors"
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"
)

// Common errors
var (
	ErrFreedList    = errors.New("operation on freed list")
	ErrEmptyList    = errors.New("operation on empty list")
	ErrAlreadyFreed = errors.New("list already freed")
)

type Node[T any] struct {
	data T
	both uintptr
}

type Block[T any] struct {
	nodes []Node[T]
	next  *Block[T]
}

type XLL[T any] struct {
	head       *Node[T]
	tail       *Node[T]
	blocks     *Block[T]
	size       int
	pinner     *runtime.Pinner
	blockSize  int
	growthRate float64
	freed      atomic.Bool
	mu         sync.RWMutex
}

type Option[T any] func(*XLL[T])

func WithBlockSize[T any](size int) Option[T] {
	return func(list *XLL[T]) {
		if size > 0 {
			list.blockSize = size
		}
	}
}

func WithGrowthRate[T any](rate float64) Option[T] {
	return func(list *XLL[T]) {
		if rate > 1.0 {
			list.growthRate = rate
		}
	}
}

func WithInitialCapacity[T any](capacity int) Option[T] {
	return func(list *XLL[T]) {
		if capacity > 0 {
			initialBlock := &Block[T]{nodes: make([]Node[T], 0, capacity)}
			list.blocks = initialBlock
			list.pinner.Pin(initialBlock)
		}
	}
}

// New function

func New[T any](options ...Option[T]) *XLL[T] {
	list := &XLL[T]{
		blockSize:  1024,
		growthRate: 2.0,
		pinner:     new(runtime.Pinner),
	}
	for _, option := range options {
		option(list)
	}
	runtime.SetFinalizer(list, (*XLL[T]).Free)
	return list
}

func (list *XLL[T]) newNode(data T) *Node[T] {
	list.mu.Lock()
	defer list.mu.Unlock()

	if list.blocks == nil || len(list.blocks.nodes) == cap(list.blocks.nodes) {
		newCapacity := list.blockSize
		if list.blocks != nil {
			newCapacity = int(float64(cap(list.blocks.nodes)) * list.growthRate)
		}
		newBlock := &Block[T]{nodes: make([]Node[T], 0, newCapacity)}
		list.pinner.Pin(newBlock)
		newBlock.next = list.blocks
		list.blocks = newBlock
	}

	node := Node[T]{data: data}
	list.blocks.nodes = append(list.blocks.nodes, node)
	list.size++
	return &list.blocks.nodes[len(list.blocks.nodes)-1]
}

func (list *XLL[T]) Free() error {
	list.mu.Lock()
	defer list.mu.Unlock()

	if !list.freed.CompareAndSwap(false, true) {
		return ErrAlreadyFreed
	}

	// Unpin all pinned objects
	list.pinner.Unpin()

	// Reset all fields
	list.head = nil
	list.tail = nil
	list.blocks = nil
	list.size = 0
	list.pinner = nil

	// Remove the finalizer
	runtime.SetFinalizer(list, nil)

	return nil
}

func (list *XLL[T]) delete(front bool) error {
	if list.IsFreed() {
		return ErrFreedList
	}
	list.mu.Lock()
	defer list.mu.Unlock()

	if list.head == nil {
		return ErrEmptyList
	}

	if list.head == list.tail {
		list.head = nil
		list.tail = nil
		list.blocks = nil
		list.size = 0
		return nil
	}

	if front {
		nextNode := (*Node[T])(unsafe.Pointer(list.head.both))
		nextNode.both = XOR(uintptr(unsafe.Pointer(list.head)), nextNode.both)
		list.head = nextNode
	} else {
		prevNode := (*Node[T])(unsafe.Pointer(list.tail.both))
		prevNode.both = XOR(prevNode.both, uintptr(unsafe.Pointer(list.tail)))
		list.tail = prevNode
	}

	if len(list.blocks.nodes) == 1 {
		list.blocks = list.blocks.next
	} else if front {
		list.blocks.nodes = list.blocks.nodes[1:]
	} else {
		list.blocks.nodes = list.blocks.nodes[:len(list.blocks.nodes)-1]
	}

	list.size--
	return nil
}

func (list *XLL[T]) traverse(f func(T), forward bool) error {
	if list.IsFreed() {
		return ErrFreedList
	}
	list.mu.RLock()
	defer list.mu.RUnlock()
	var prev uintptr
	var curr unsafe.Pointer
	if forward {
		curr = unsafe.Pointer(list.head)
	} else {
		curr = unsafe.Pointer(list.tail)
	}
	for curr != nil {
		node := (*Node[T])(curr)
		f(node.data)
		next := XOR(prev, node.both)
		prev = uintptr(curr)
		// Conversion of a Pointer to an uintptr and back, with arithmetic.
		// Explicitly following pattern (3)
		curr = unsafe.Pointer(next)
	}
	return nil
}

func (list *XLL[T]) insert(data T, front bool) error {
	if list.IsFreed() {
		return ErrFreedList
	}
	newNode := list.newNode(data)
	list.mu.Lock()
	defer list.mu.Unlock()

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else if front {
		newNode.both = uintptr(unsafe.Pointer(list.head))
		list.head.both = XOR(list.head.both, uintptr(unsafe.Pointer(newNode)))
		list.head = newNode
	} else {
		newNode.both = uintptr(unsafe.Pointer(list.tail))
		list.tail.both = XOR(list.tail.both, uintptr(unsafe.Pointer(newNode)))
		list.tail = newNode
	}
	return nil
}

func (list *XLL[T]) DeleteFront() error {
	return list.delete(true)
}

func (list *XLL[T]) DeleteBack() error {
	return list.delete(false)
}

func (list *XLL[T]) TraverseForward(f func(T)) error {
	return list.traverse(f, true)
}

func (list *XLL[T]) TraverseBackward(f func(T)) error {
	return list.traverse(f, false)
}

func (list *XLL[T]) InsertFront(data T) error {
	return list.insert(data, true)
}

func (list *XLL[T]) InsertBack(data T) error {
	return list.insert(data, false)
}
