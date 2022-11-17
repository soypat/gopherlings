// Problem:
// TODO

// I AM STILL GOING

package main

import "fmt"

func main() {
	// Huh? Someone is trying to give this movie
	// a 10/10 score but their efforts are not succeeding.
	// Help them out by modifying the setMovie function.
	var myMovie movie
	setMovie(myMovie)
	fmt.Println(myMovie)
}

type movie struct {
	title string
	score int
}

// setMovie for all intents and purposes creates a movie from scratch.
// We should probably change it so it returns a movie type instead... and
// probably change the name to something that describes it better.
func setMovie(m movie) {
	// Dot indexing is how we access fields in Go.
	m.title = "Everything Everywhere All At Once"
	m.score = 10
	// The movie argument is a copy, so all changes within this function
	// are lost on function exit. We should return a value here.
}
