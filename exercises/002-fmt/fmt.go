//
// Problem:
// This program prints out `hello world` but as of now throws a compiler error
// claiming `fmt` is undefined. This is because it is missing the import statement.
//
// The import statement starts with the `import` keyword and is followed by
// the name of the import wrapped in double quotes.
//
//  import "fmt"
//
// For when there's more than a single import statement the import names are
// listed inside parentheses.
//
//  import (
//	   "fmt"
//     "strings"
//  )
//
// It is rare to modify the import statements since most IDE's can automatically
// deduce which import you are referring to and automatically add it on save.

// I AM STILL GOING

package main

// The import statement always follows the `package`
// declaration.

func main() {
	fmt.Println("hello world")
}
