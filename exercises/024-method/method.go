// Problem:
// TODO

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

// Worth noting we could also solve this excercise by simply renaming this method
// to "String". The fmt package checks all arguments for the `String` method and
// uses it to print out the arguments, overriding the default printing algorithm.
func (b book) Info() string {
	yearString := strconv.Itoa(b.year) // convert year to string!
	return "\"" + b.title + "\" by " + b.author + " (" + yearString + ")"
}
