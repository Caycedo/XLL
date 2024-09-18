package main

import (
	"fmt"
	"github.com/Caycedo/XLL"
)

func main() {
	// Create a new XLL with custom options
	list := XLL.New[int](
		XLL.WithBlockSize[int](64),
		XLL.WithGrowthRate[int](1.5),
		XLL.WithInitialCapacity[int](32),
	)

	// Insert elements at the front and back
	fmt.Println("Inserting elements:")
	for i := 0; i < 5; i++ {
		if err := list.InsertFront(i); err != nil {
			fmt.Printf("Error inserting front: %v\n", err)
		}
		if err := list.InsertBack(i + 10); err != nil {
			fmt.Printf("Error inserting back: %v\n", err)
		}
	}

	// Print elements forward
	fmt.Println("Forward print:")
	if err := list.PrintForward(); err != nil {
		fmt.Printf("Error printing forward: %v\n", err)
	}

	// Print elements backward
	fmt.Println("Backward print:")
	if err := list.PrintBackward(); err != nil {
		fmt.Printf("Error printing backward: %v\n", err)
	}

	// Delete from front and back
	fmt.Println("Deleting elements:")
	if err := list.DeleteFront(); err != nil {
		fmt.Printf("Error deleting front: %v\n", err)
	}
	if err := list.DeleteBack(); err != nil {
		fmt.Printf("Error deleting back: %v\n", err)
	}

	// Print after deletion
	fmt.Println("After deletion:")
	if err := list.PrintForward(); err != nil {
		fmt.Printf("Error printing forward: %v\n", err)
	}

	// Demonstrate the iterator
	fmt.Println("Using iterator:")
	for it := list.Iterator(); it.Next(); {
		fmt.Printf("%v ", it.Value())
	}
	fmt.Println()

	// Free the list
	if err := list.Free(); err != nil {
		fmt.Printf("Error freeing list: %v\n", err)
	}

	// Attempt to use the list after freeing
	fmt.Println("Attempting to insert after freeing:")
	if err := list.InsertFront(100); err != nil {
		fmt.Printf("Error inserting after free: %v\n", err)
	}

	// Create a new list to demonstrate error on freed list
	newList := XLL.New[int]()
	if err := newList.InsertFront(1); err != nil {
		fmt.Printf("Error inserting into new list: %v\n", err)
	}
	if err := newList.Free(); err != nil {
		fmt.Printf("Error freeing new list: %v\n", err)
	}

	// This should return an error, not panic
	if err := newList.PrintForward(); err != nil {
		fmt.Printf("Error printing freed list: %v\n", err)
	}
}
