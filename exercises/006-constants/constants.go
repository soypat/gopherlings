// Problem:
// In Go constants are just what the name suggests, constant- unchangeable,
// immutable.
// Go supports constant string, numeric, boolean values.
//
// You can't modify or reassign to a constant in Go.
package main

import "fmt"

func main() {
	const abc = "abc"
	fmt.Println(abc)

	const sum = abc + "tyk"
	fmt.Println(sum)

	// Print out the declared value of sum with "cheeky addition" appended to end.
	sum = sum + "cheeky addition"
	fmt.Println(sum)

	// In Go variable names must start with a letter.
	const 20number = 20
	fmt.Println(math.Sin(20number)) 

}
