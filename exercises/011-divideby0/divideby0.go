//
// Problem:
// Dividing an integer by zero is an undefined operation. To deal with
// this Go will `panic` to avoid undefined behaviour. Panics result in
// crashes if not caught by a recover(), more on that later.

package main

import "fmt"

func main() {
	var z int
	// Add an if statement that sets z
	// to 7 if z is equal to zero.
	fmt.Println(49/z + 7)
}
