//
// Problem:
// One of the most interesting interfaces in Go is the `image.Image` interface.
//
// It is a curiously simple interface but broad in scope. In a few lines
// of code we may define an entire image.

// I AM STILL GOING

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
)

const imageSize = 100

func main() {
	// You may have not yet seen files, this is how a file is created in Go.
	// File creation can fail so an error is returned.
	fp, err := os.Create("stars.png")
	if err != nil {
		// if an error is returned, we end the program with a call to log.Fatal
		// so that the error is printed.
		log.Fatal(err)
	}
	// Right now the sky is pretty dark. Let's fill it with stars.
	// By setting the value of s (stars type) we'll change the probability
	// of there being a star on any given pixel. Try setting it to a value
	// between 0 and 100 and run the code again.
	var s stars
	err = png.Encode(fp, s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("written a starlit sky to stars.png")
}

type stars int

func (s stars) At(i, j int) color.Color {
	// To determine if pixel is a star, we see if
	// a random integer between 0 and 100 is lesser than
	// the value of the stars receiver.
	if rand.Intn(imageSize*imageSize) < int(s) {
		return color.White
	}
	return color.Black
}

// ColorModel returns the range of colors that can be expected from the At method.
// Since we are drawing a black and white sky a color.GrayModel is enough for this use case.
func (s stars) ColorModel() color.Model {
	return color.GrayModel
}

// Bounds defines a 100x100 grid of starlit sky.
func (s stars) Bounds() image.Rectangle {
	return image.Rect(0, 0, imageSize, imageSize)
}
