// Problem:
// Booleans have the following logical operators defined for them:
//
//  &&    conditional AND    p && q  results true only if both p and q are true.
//  ||    conditional OR     p || q  results true if either p or q are true.
//  !     NOT                !p      results true if p is false.
//

package main

import "fmt"

func main() {
	a := true
	b := false
	// Fix the following line so that
	// it prints out the result of a OR b.
	fmt.Println(a + b)

	// Fix the following line so that
	// it prints out the result of a AND b.
	fmt.Println(a & b)
}
