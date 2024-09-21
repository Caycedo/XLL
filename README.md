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

XLL offers comparable performance to standard doubly linked lists for most operations, with the added benefit of reduced memory usage. Here are the benchmark results:

### XLL Benchmark Results

This benchmark compares the performance of XLL (XOR Linked List) against standard slices and doubly linked lists in Go.

#### System Information
- **OS**: Windows
- **Architecture**: amd64
- **CPU**: 13th Gen Intel(R) Core(TM) i9-13900H

#### Operation Benchmarks

| Operation | XLL | Slice | Doubly Linked List |
|-----------|-----|-------|---------------------|
| Insert Front | 42.09 ns/op | - | 54.95 ns/op |
| Insert Back | 42.48 ns/op | 10.30 ns/op | - |
| Delete Front | 20.54 ns/op | - | - |
| Delete Back | 20.41 ns/op | - | - |
| Traverse Forward | 1463252 ns/op | 96078 ns/op | 3189529 ns/op |
| Traverse Backward | 1536292 ns/op | - | - |

#### Memory Usage

| Data Structure | Memory Usage |
|----------------|--------------|
| XLL | 16,777,408 B |
| Slice | 40,280,064 B |
| Doubly Linked List | 55,997,984 B |

#### Additional XLL Benchmarks

| Operation | Performance |
|-----------|-------------|
| Insert Front | 45.76 ns/op |
| Insert Back | 43.70 ns/op |
| Insert Front (with options) | 41.27 ns/op |
| Insert Back (with options) | 41.45 ns/op |

#### Analysis

1. **Insertion**: XLL performs competitively, with insert operations ranging from 41-46 ns/op. It outperforms the doubly linked list (54.95 ns/op) but is slower than slice insertion (10.30 ns/op).

2. **Deletion**: XLL shows excellent performance in deletion operations, consistently around 20.4-20.5 ns/op for both front and back deletions.

3. **Traversal**: XLL (1,463,252 ns/op) is significantly faster than the doubly linked list (3,189,529 ns/op) but slower than slice traversal (96,078 ns/op).

4. **Memory Usage**: XLL shows superior memory efficiency, using only 16,777,408 bytes compared to 40,280,064 bytes for slices and 55,997,984 bytes for doubly linked lists.

In summary, XLL offers a balanced performance profile with competitive insertion and deletion speeds, faster traversal than traditional linked lists, and significantly lower memory usage compared to both slices and doubly linked lists.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
