# XLL: XOR Linked List Implementation in Go

XLL is an implementation of an XOR Linked List in Go. This data structure uses the XOR operation to combine the addresses of the previous and next nodes in a single field, offering bidirectional traversal capability while using less memory than a traditional doubly linked list.

## Features

- Memory-efficient bidirectional list using XOR linking
- Supports traversal in both directions with only one address field per node
- Generic type support
- Thread-safe operations
- Customizable block size and growth rate
- Efficient insertion and deletion at both ends
- Iterator support

## How It Works

In an XOR Linked List, each node stores a single pointer, which is the XOR of the addresses of the previous and next nodes. This clever use of the XOR operation allows for bidirectional traversal while using only one pointer field per node, potentially reducing memory usage compared to a traditional doubly linked list.

To traverse the list, we XOR the address of the current node with the previous (or next) node's address to get the address of the next (or previous) node. This technique allows us to move both forward and backward through the list.

## Installation

To use XLL in your Go project, you can install it using `go get`:

```bash
go get github.com/Caycedo/XLL
```

## API Reference

- `New[T any](options ...Option[T]) *XLL[T]`: Create a new XLL
- `InsertFront(data T) error`: Insert an element at the front
- `InsertBack(data T) error`: Insert an element at the back
- `DeleteFront() error`: Delete the front element
- `DeleteBack() error`: Delete the back element
- `TraverseForward(f func(T))`: Traverse the list from front to back
- `TraverseBackward(f func(T))`: Traverse the list from back to front
- `PrintForward()`: Print the list from front to back
- `PrintBackward()`: Print the list from back to front
- `Free()`: Free the list and its resources

## Customization

You can customize the XLL behavior using options:

```go
list := XLL.New[int](
    XLL.WithBlockSize[int](64),
    XLL.WithGrowthRate[int](1.5),
    XLL.WithInitialCapacity[int](32),
)
```

## Performance

XLL offers comparable performance to standard doubly linked lists for most operations, with the added benefit of reduced memory usage. Benchmark results are available in the `benchmark_test.go` file.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
