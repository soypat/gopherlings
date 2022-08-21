// Problem:
// https://gobyexample.com/pointers
// Go supports pointers, allowing to pass values by reference
// to functions and methods.
//
// What this means is that a function with a pointer argument
// will be able to modify the value the pointer points to.
// To obtain the reference to a value in Go we use the "&" operator:
//
//  var a int = 2
//  ref := &a
//
// ref will be of type *int, which is to say ref is a reference to an
// integer. To access the underlying value we must dereference the pointer.
// using the * operator, which precedes the pointer value:
//
//  fmt.Println(*ref) // prints 2
//  fmt.Println(ref)  // prints the location of the pointer in memory
//
// To assign to the pointer we also dereference:
//
//  *ref = 1337
//  fmt.Println(*ref) // prints 1337
//

// I AM STILL GOING

package main

import "fmt"

func main() {
	// This is the hardest exercise so far, so be ready for a challenge!
	// We need to modify at least 3 lines in this program to achieve
	// the desired result: zeroIt should modify value without needing
	// a return parameter.
	value := 2
	zeroIt(value) // We may need to pass `value` by reference here.
	fmt.Println(value)
}

// We want to modify p from inside a function- so we
// need the help of a pointer! Change the argument type
// so that zeroIt accepts a pointer and then modify the function
// body so that we assign zero to p.
func zeroIt(p int) {
	p = 0
}
