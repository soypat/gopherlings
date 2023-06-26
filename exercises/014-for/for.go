//
// Problem:
// For statements can also have init and post statements in the clause.
//
//  for i := 0; i <= 10; i = i + 1 {
//  	fmt.Println(i)
//  }
//
// The above for loop will iterate from 0 to 10.
// The following single-condition for statement
// will perform the same operation:
//
//  i := 0
//  for i <= 10 {
//  	fmt.Println(i)
//  	i = i + 1
//  }
//

// I AM STILL GOING

package main

func main() {
	// Snap, this for syntax is hard to remember!
	// It's very common for gophers to have to look up
	// how to write for loops long after they've seen the
	// syntax for the first time. Keep in mind when fixing
	// this for statement that
	//  1. The first statement is the init (runs once at start)
	//  2. Second is the condition. For loop executes while condition yields true
	//  3. Third is the post statement. This runs at the end of each loop.
	for i <= 3; i := 0;  i = i + 1 {
		fmt.Println(i)
	}
}
