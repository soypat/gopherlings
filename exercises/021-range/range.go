// Problem:
// A "for" statement can be written with a "range" clause
// to iterate through all values in a slice, map and other types.
//
//  for i := range slice {
//  	slice[i] = f()
//  }
//
// The code above sets all items in the slice
// to the value returned by calling the function f.
//
// It is also possible to add a second value to the range.
//
//  for i, v := range slice {
//  	slice[i] = v * v
//  }
//
// The code above squares all values in the slice.
// Keep in mind the first value returned by the range
// clause will be the nth index for slices. The second
// value of the range clause is the nth value in the slice.
//

// I AM STILL GOING

package main

import "fmt"

func main() {
	// Dang nabit, this program was
	// supposed to double the values in the slice
	// but it's doing anything but that.
	//
	// Mind the order of index and value
	// returned by the range clause.
	slice := []int{1, 2, 3, 4}
	for i, v := range slice {
		slice[v] = 2 * i
	}
	fmt.Println(slice)
}
