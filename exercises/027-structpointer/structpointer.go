// Problem:
// TODO

// I AM STILL GOING

package main

import "fmt"

func main() {
	// Same program we saw before, but now
	// we want to solve it using pointers.
	// If we have a pointer to a struct, modifying
	// the fields of the struct will be reflected
	// in the struct the pointer references.
	var myMovie movie
	setMovie(myMovie)
	fmt.Println(myMovie)
}

type movie struct {
	title string
	score int
}

// You may need to only modify one line here.
func setMovie(m movie) {
	m.title = "Blue Planet"
	m.score = 10
}
