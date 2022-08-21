// Problem:
// TODO

// I AM STILL GOING

package main

import (
	"fmt"
	"strings"
)

func main() {
	// One must expect anything from user input.
	// Us humans are like that; unpredictable and
	// not fond of compiler errors.
	// It is common to "sanitize" user input before
	// processing or displaying it. Here we made the mistake
	// of trying to call PrintSafe on a regular string- but PrintSafe
	// takes a `safe` type as an argument which is really just
	// a string with it's own distinct type.
	// This prevented us from printing unsanitized input! To sanitize
	// call `sanitize` on unsafeInput to obtain a "safe" string!
	unsafeInput := "stupid compiler"
	PrintSafe(unsafeInput)
}

type safe string

func PrintSafe(s safe) {
	fmt.Println(s)
}

func sanitize(possibleBadString string) safe {
	safeString := strings.ReplaceAll(possibleBadString, "stupid", "****")
	return safe(safeString)
}
