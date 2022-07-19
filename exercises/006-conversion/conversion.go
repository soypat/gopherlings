// Problem:
// TODO
// https://go.dev/tour/basics/13

package main

import "fmt"

func main() {
	var three int = 3
	var oneThird float64 = 1.0 / 3.0
	var twoPlusThree int = 2 + three
	// These two variables are of distint type. To perform operations
	// between variables they must be of the same type! Convert one of
	// the variables on this line so that the result of 1/3+2+3 is printed out.
	// Don't add lines or change the lines above!
	fmt.Println(oneThird + twoPlusThree)
}
