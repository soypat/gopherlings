// Problem:
// TODO
//

// I AM STILL GOING

package main

import "fmt"

func main() {
	// Notice that unlike with composite literals like
	// structs, maps and slices we may not initialize
	// a XORer with curly braces {}, we actually perform a type
	// conversion from the bool primitive type to XORer.
	// This works since XORer is of underlying type bool.
	True := XORer(true)
	False := XORer(false)
	fmt.Println("true ^ true =", True.xor(true))
	fmt.Println("false ^ true =", False.xor(true))
	fmt.Println("false ^ false =", False.xor(false))
}

// It comes as a surprise to many newcomers to Go that you may define
// methods on any user defined type, including types that are composed
// of an underlying bool type, like in this case!
type XORer bool

// There is no XOR operation in Go, instead we use the inequality
// operator to perform a xor between two values of the bool type.
// This code has an error due to operations between two distinct types.
// These two values should be of same type to perform the inequality operation!
func (a XORer) xor(b bool) bool {
	return a != b
}
