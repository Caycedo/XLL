package XLL

import "unsafe"

// Iterator represents an iterator over the XOR linked list elements.
type Iterator[T any] struct {
	current *Node[T]
	prev    uintptr
	list    *XLL[T]
}

// Next advances the iterator and returns whether there is a next element.
func (it *Iterator[T]) Next() bool {
	if it.current == nil {
		return false
	}
	next := XOR(it.prev, it.current.both)
	it.prev = uintptr(unsafe.Pointer(it.current))
	it.current = (*Node[T])(unsafe.Pointer(next))
	return it.current != nil
}

// Value returns the current value of the iterator.
// It panics if called when there is no current value.
func (it *Iterator[T]) Value() T {
	if it.current == nil {
		panic("Value called on exhausted iterator")
	}
	return it.current.data
}

// Iterator returns an iterator for the list that can be used with range.
func (list *XLL[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{
		current: list.head,
		prev:    0,
		list:    list,
	}
}
