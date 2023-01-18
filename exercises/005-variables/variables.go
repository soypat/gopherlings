//
// https://gobyexample.com/variables
//
// Problem:
// In Go variables are declared with their type so that the compiler
// can check correctness of function calls.
//
// The keyword `var` allows for declaring one or more variables
// The following code declares a variable named `a`, of type `int`.
//
//  var a int
//
// One can then assign to the newly created value below the declaration as
// many times as you'd like:
//
//  a = 2
//
// You can also declare and assign in the same line having the same effect:
//
// var a int = 2
//
// To declare multiple variables in one line
//
//  var a, b string = "hello", "world"
//
// In Go it is common to omit the type in a variable declaration:
//
//  var a = "look ma, no type!"
//
// With the shortest (and possibly most common) way to declare and initialize
// a variable being the shorthand `:=`
//
//  a := "as short as you can"
//  b := 4
//
// In Go, variables that are declared without initialization are "zero-valued"
// For example, strings will be the empty string "", numbers will be 0.
//

// I AM STILL GOING

package main

import "fmt"

func main() {
	// The value being assigned to `a` has a type
	// that does not correspond to `a`'s type.
	// Fix this line so that the following print
	// prints "123".
	var a string

	a = 123
	fmt.Println(a)

	// We are trying to declare and initialize variables `b` and `c`
	// but there's an syntax error. Adding just 2 characters should suffice to it!
	var b c int = 2*1 2*2
	fmt.Println(b, c)

	// Someone tried using the shorthand initialization but got it wrong.
	d = "pear"
	fmt.Println(d)
}
