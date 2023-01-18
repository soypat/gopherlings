//
// Problem:
// Go does not have classes. However, you can define methods on types.
//
// A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// In this example, the Info method has a receiver of type book named b.
//
// And remember, methods are just functions. The method below could be rewritten to be
// a regular function with no change in functionality:
//
//  func Info(b book) string {
//		// same code
//  }

// I AM STILL GOING

package main

import (
	"fmt"
	"strconv"
)

func main() {
	checkedOut := book{
		title:  "Los Hijos de los Días",
		author: "Eduardo Galeano",
		year:   2012,
	}
	// Well this program does something remotely correct
	// at least... it prints out the unformatted contents of the struct.
	// We want to print formatted information on the book though- we
	// could do this by calling a certain method that returns formatted
	// information on the book.
	fmt.Println(checkedOut)
}

type book struct {
	title  string
	author string
	year   int
}

// Worth noting we could also solve this exercise by simply renaming this method
// to "String". The fmt package checks all arguments for the `String` method and
// uses it to print out the arguments, overriding the default printing algorithm.
// See the fmt.Stringer type.
func (b book) Info() string {
	yearString := strconv.Itoa(b.year) // convert year to string!
	return "\"" + b.title + "\" by " + b.author + " (" + yearString + ")"
}
