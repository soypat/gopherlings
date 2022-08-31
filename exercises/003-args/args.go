//
// Problem:
// Below is a program that used to prints out values. Looks like
// whomever refactored this was not aware fmt.Println accepts multiple arguments
// with the help of commas to separate arguments of different types
//
//  fmt.Println("hello", "world", 99)
//
// The previous program prints out "hello world 99". Variadic functions accept
// an arbitrary number of arguments.

// I AM STILL GOING

package main

import "fmt"

func main() {
	// Knowing how to separate arguments in function calls fix
	// the following line so that it prints "go 2 true"
	fmt.Println("go" 2 true)
}
