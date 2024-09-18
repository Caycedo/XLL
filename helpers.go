package XLL

import "fmt"

func XOR(a, b uintptr) uintptr {
	return a ^ b
}

func (list *XLL[T]) Size() int {
	list.mu.RLock()
	defer list.mu.RUnlock()
	return list.size
}

func (list *XLL[T]) IsFreed() bool {
	return list.freed.Load()
}

func (list *XLL[T]) PrintForward() error {
	err := list.TraverseForward(func(data T) {
		fmt.Printf("%v ", data)
	})
	if err != nil {
		return err
	}
	fmt.Println()
	return nil
}

func (list *XLL[T]) PrintBackward() error {
	err := list.TraverseBackward(func(data T) {
		fmt.Printf("%v ", data)
	})
	if err != nil {
		return err
	}
	fmt.Println()
	return nil
}
