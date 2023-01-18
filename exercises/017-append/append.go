// Problem:
// append is a built-in that helps construct slices whose size is undefined
// at runtime.
//
// To append an item to an existing slice one uses the append function
//  short := []string{"first"}
//  longer := append(short, "second")
//
// The result is stored to `longer` and is a slice of two elements, "first" followed
// by "second".
//
// A common pattern is to append to the same slice
//
//  elements = append(elements, "an item")
//

// I AM STILL GOING

package main

import "fmt"

func main() {
	var evens []int
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			// Knowing the basic mechanics of slices, how
			// would you go about constructing a slice of
			// all even numbers from 1 up to 20?
			evens = ??
		}
	}
	fmt.Println(evens)
}
