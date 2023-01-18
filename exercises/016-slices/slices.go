//
// Problem:
// A slice is the most common way to store ordered, consecutive elements
// of the same type in a single structure. In other languages a slice is known
// as an array or even lists.
//
// To initialize a slice one uses the make built-in:
//
//  desiredLength := 23
//  s := make([]int, desiredLength)
//
// Slices are zero-indexed in Go:
//
//  A[0] // access to A's first element
//  A[9] // access to A's tenth element
//
// Go does have an array type but these see very limited use since an array's length
// is also part of it's type.
//

// I AM STILL GOING

package main

func main() {
	// How long should the slice be to house all values?
	beatles := make([]string, ??)

	beatles[0] = "john"
	beatles[1] = "george"
	beatles[2] = "ringo"
	beatles[3] = "paul"

	// Change this line so it prints the drummer's name.
	fmt.Println(beatles[??])
}
