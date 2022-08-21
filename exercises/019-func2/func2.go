//
// Problem:
// Functions may have multiple arguments and multiple return values.
// To create a function with multiple return values surround the
// return types in parentheses.
//
//  func fruits() (string, string) {
//  	return "banana", "tangerine"
//  }
//
// If two or more consecutive arguments have the same type one can omit
// the type name until the last argument.
//
//  func divmod(a, b int) (div, mod int) {
//  	div = a/b
//  	mod = a%b
//  	return div, mod
//  }
//

// I AM STILL GOING

package main

import "fmt"

func main() {
	fmt.Println(return1And2())
}

// This function should return two values, 1 and 2.
// Fix the function signature (the line with the `func` keyword)
// by adding the two returned parameters with names like
// the last example in the problem description.
func return1And2() {
	a = 1
	b = 2
	return a, b
}
