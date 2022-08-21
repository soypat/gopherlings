// Problem:
// Map types are Go's associative data type commonly referred to as
// hash or dicts in other languages.
//
// A map contains an unordered group of elements of one type, called the element
// type, indexed by a set of unique keys of another type.
//
// In Go a map must be initialized before use. To do this we use the
// make built-in function:
//
//  myMap := make(map[int]int)
//
// When searching a map for a non-existent value the element's zero-value
// will be returned.

// I AM STILL GOING

package main

import "fmt"

func main() {
	// A map with string keys and integer elements.
	var age map[string]int

	// The program will crash if the map is not initialized correctly
	// before accessing a key.
	// Fix the declaration above!
	age["paul"] = 32
	age["miguela"] = 27
	age["sarah"] = 78

	fmt.Println(age)

	// When accessing a map be careful with
	// the key- it must be an exact match of the
	// key you are searching for.
	// If the key is not found then the access
	// yields the zero value for the element type.
	fmt.Println("Sarah's Age:", age["Sarah"])
}
