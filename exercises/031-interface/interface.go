//
// Problem:
// An interface type is defined as a set of method signatures.
//
// A value of interface type can hold any value that implements those methods.
//
// When a type implements the methods contained in the interface it is said that
// the type "implements" that interface.
//
// The most common interfaces in Go are the io.Reader and io.Writer interfaces.

// I AM STILL GOING

package main

import "fmt"

func main() {
	const myNumber = 8
	// Your mission is to implement the Decimal type and to implement the
	// Number interface so that these lines do not yield error. Take note!
	// You do not have to modify the rest of the code in the main() function!
	// These lines below once uncommented should run with no problem once the
	// problem at hand has been solved.

	// d := Decimal(myNumber)
	// PrintNumber(d)

	b := Binary(myNumber)
	PrintNumber(b)

	r := Roman(myNumber)
	PrintNumber(r)
}

type Number interface {
	Numeral() string
}

func PrintNumber(n Number) {
	fmt.Println(n.Numeral())
}

// The Binary and Roman types implement the Number interface *implicitly*.
// In contrast with other languages, in Go there is no "implements" clause. One must
// only write the methods and the compiler will know that the methods have been implemented.

type Binary int
type Roman int

func (b Binary) Numeral() string {
	return fmt.Sprintf("%b", b)
}

func (r Roman) Numeral() (numeral string) {
	switch r {
	case 0:
		numeral = "0"
	case 1:
		numeral = "I"
	case 2:
		numeral = "II"
	case 3:
		numeral = "III"
	case 4:
		numeral = "IV"
	case 5:
		numeral = "V"
	case 6:
		numeral = "VI"
	case 7:
		numeral = "VII"
	case 8:
		numeral = "VIII"
	case 9:
		numeral = "IX"
	case 10:
		numeral = "X"
	default:
		numeral = "error: roman numeral too large"
	}
	return numeral
}
