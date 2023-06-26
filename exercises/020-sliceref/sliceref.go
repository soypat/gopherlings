//
// Problem:
// Up to now we have seen functions receive a copy of their arguments
// and are unable to modify them directly. This is not the case for all types in Go.
//
// The slice type we saw just now actually holds a reference to the underlying
// memory contained in the slice. What this means is that a copy of a slice
// does not copy the memory in the indices of the slice. This can lead
// to unexpected behaviour if gophers are not careful.
//
// This reference behaviour may be clearer after learning about pointers.

// I AM STILL GOING

package main

import "fmt"

func main() {
	fibonacci := []int{0, 1, 1, 2, 3, 5, 8, 13}
	fiboAdded := fibonacci
	for i := range fiboAdded {
		fiboAdded[i] += i
	}
	// Oof. We have inadvertently modified both slices
	// Try using the clone function below to make a
	// deep copy of the fibonacci slice so that edits to
	// one are not reflected in the other.
	fmt.Println(fibonacci, fiboAdded)
}

func clone(a []int) []int {
	return append([]int{}, a...)
}
