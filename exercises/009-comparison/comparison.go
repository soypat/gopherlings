//
// Problem:
// Go has comparison operators which compare two values and yield a boolean value.
//
// Go has support for the following comparison operators:
//     a == b   means "a equals b"
//     a < b    means "a is less than b"
//     a > b    means "a is greater than b"
//     a != b   means "a does not equal b"
//     a <= b    means "a is less-or-equal than b"
//     a >= b    means "a is greater-or-equal than b"

// I AM STILL GOING

package main

import "fmt"

func main() {
	number := 4
	// % is the modulo operator. It returns the remainder of the division.
	// A number is even when the remainder of number/2 is zero.
	remainder := number % 2


	var is4even bool
	// Something is wrong with this comparison.
	is4even = remainder = 0

	// This should print "4 is even:" followed by true or false depending
	// on whether 4 (the number variable) is really a even number.
	fmt.Println(number, "is even:", is4even)
}
