// Problem:
// `if` statements specify a block of code to execute given a condition is true.
// These come in a variety of flavours.
//
// Simple if statements:
//
//  if x > max {
//  	max = x
//  }
//
// Go has support for the following comparison operators:
//     a == b   means "a equals b"
//     a < b    means "a is less than b"
//     a > b    means "a is greater than b"
//     a != b   means "a does not equal b"
//     a <= b    means "a is less-or-equal than b"
//     a >= b    means "a is greater-or-equal than b"

package main

import "fmt"

func main() {
	const tuesdayDiscount = 0.50 // 50% discount! Wow!
	isTuesday := 1

	price := 32.00
	// fix this condition to accept numbers:
	if isTuesday {
		price = price * (1 - tuesdayDiscount)
	}
	fmt.Println("The price is", price)
}
