// Problem:
// Go is a statically, strongly typed language. This means
//  - All values in go have a defined type at compile time (static typing)
//  - Operations between values of different types is not allowed
//
// Below is a program that used to prints out values. Looks like
// whomever refactored this was not aware fmt.Println accepts multiple arguments
// with the help of commas to separate arguments of different types
//
//  fmt.Println("hello", "world", 99) // Prints out "hello world 99".
//
// Beware divide by zeros: they will crash your program. Whenever Go can help
// it will detect divide by zeros at compile time and warn you.
//
// Lastly, since go is strongly typed it does not allow operations between
// different types.
//
// For example: You can't add a boolean to a number without converting it
// via a function.
//
// The operations available for each type depend on the type. 
//
// Booleans have the following logical operators defined for them:
//
//  &&    conditional AND    p && q  results true only if both p and q are true.
//  ||    conditional OR     p || q  results true if either p or q are true.
//  !     NOT                !p      results true if p is false.

package main

import "fmt"

func main() {
	fmt.Println("go" 2)   // Prints "go 2"

	fmt.Println(49/0 + 7) // Prints out the result of 49 divided by (0+7)
	
	fmt.Println(3.14159 * true) // Prints out 3.14159 times 1

	// Prints out the result of boolean and operation between true and false:
	fmt.Println(true & false)
	fmt.Println(true + true) // Prints out "true true"

	// % is the modulo operator. It returns the remainder of the division.
	// This should print "4 is even:" followed by true or false depending
	// on whether 4 is really a even number.
	remainder := 4%2
	is4even := remainder = 0 // Fix this conditional!
	fmt.Println("4 is even:", is4even)
}
