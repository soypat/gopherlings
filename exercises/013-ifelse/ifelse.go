// Problem:
// Go also has support for if-else statements
//
// if statements with an else branch that executes if the condition is false:
//
//  if x > max {
//  	max = x
//  } else {
//  	x = max
//  }
//
// if statements with else if and/or a preceding statement
//
//  if x := f(); x > max {
//  	max = x
//  } else if x == max {
//  	max = max + 1
//  } else {
//  	x = max
//  }
//

// I AM STILL GOING

package main

import "fmt"

func main() {
	number := 7
	if number > 0 {
		fmt.Println("number is greater than 0")
	} {
		fmt.Println("number is less than 0")
	}

	// Please make sure conditionals are correct.
	if number = 0 {
		fmt.Println("number is zero")
	} else if number % 2 != 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}
}
