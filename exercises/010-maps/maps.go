// Problem:
// Map types are Go's associative data type commonly referred to as
// hash or dicts in other languages.
//
// A map contains an unordered group of elements of one type, called the element
// type.
//
// In Go a map must be initialized before being usable. To do this we use the
// make built-in function:
//
//  myMap := make(map[int]int)
//

package main

import "fmt"

func main() {
	var age map[string]int

	// The program will panic if the map is not initialized correctly
	// Fix the declaration above!
	age["paul"] = 32
	age["miguela"] = 27
	age["sarah"] = 78

	fmt.Println(age)

	// Sarah's age should display here. Keep in mind
	// map finds exact matches
	fmt.Println(age["Sarah"])
}
