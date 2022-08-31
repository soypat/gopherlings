// Problem:
// STOP! Don't run this code yet. There's a fatal flaw in it.
// This code will run **forever** unless it's fixed.
//
// In Go `for` is the only looping control structure. There are however a few
// ways of writing them. The one shown below is the
// "for statement with single condition". It will run the code inside the `{ }`
// repeatedly as long as the boolean condition evaluates to true.
//
// In the code below i is initialized to 0, which then causes the for's
// condition to evaluate to true. Since `i` is never changed the code is run
// forever.

// I AM STILL GOING

package main

import "fmt"

func main() {
	i := 0
	for i <= 3 {
		fmt.Println(i) // Prints out numbers 0 to 3
		// Did we forget to add 1 to something?
	}
}
