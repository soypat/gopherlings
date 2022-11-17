// Problem:
// Just like functions can have pointer arguments, methods can have pointer receivers.
//
// There are two reasons to use a pointer receiver.
//
// The first is so that the method can modify the value that its receiver points to.
//
// The second is to avoid copying the value on each method call.
// This can be more efficient if the receiver is a large struct, for example.
//
// In general, all methods on a given type should have either value or pointer receivers,
// but not a mixture of both. We'll see why soon.

// I AM STILL GOING

package main

import "fmt"

func main() {
	var id universalID
	id.Set(20)
	// It seems we're printing the zero value of universalID (0).
	// Somehow, we'd want id.Set to have the desired effect of setting
	// id to the argument of Set, we can do this by changing the code in
	// the Set method.
	fmt.Println(id.Get())
}

type universalID int

// Set sets the receiver universalID to newID.
func (uid universalID) Set(newID int) {
	uid = universalID(newID)
}

func (uid universalID) Get() int {
	return int(uid)
}
