//
// Problem:
// Functions are central to execution of a Go program.
// You may have noticed you've been writing your Go programs
// inside the body of a function called `main`.
//
//  func main() {
//  	// Your code here!
//  }
//
// The `main` function of the `main` package defines an executable Go program.
// It is called automatically when the program is run.
//
// You can also define executable pieces of code outside of `main`!
// Here's a function that adds two integers together and returns the result.
//
//  func add(a int, b int) int {
//  	c := a + b
//  	return c
//  }
//
// This function can then be used in the rest of your code easily!
//
//  sum := add(1, 2)
//
//

// I AM STILL GOING

package main

import "fmt"

func main() {
	sq := square(3)
	fmt.Println(sq)
}

// Oops! Someone wanted to square a number
// and return the result but it looks like
// something is off.
// Remember that the type that follows the parentheses
// in a function is the return type. When a function is
// defined with a return value one must also remember to
// specify which value to return using the `return` keyword.
func square(a int) int {
	a * a
}
