// Problem:
// TODO

// I AM STILL GOING

package main

import "fmt"

func main() {
	fruitScores := map[string]int{
		"apple":     5,
		"banana":    8,
		"tangerine": 10,
	}
	// Why won't this compile!?
	// Go is picky on unused variables.
	// In this particular case we can reduce
	// the overall code by quite a lot. Use
	// v inside this for block to replace
	// map accesses and print out fruits scoring
	// higher than 7.
	for k, v := range fruitScores {
		if fruitScores[k] > 7 {
			fmt.Println(k, " has score ", fruitScores[k])
		}
	}
}
