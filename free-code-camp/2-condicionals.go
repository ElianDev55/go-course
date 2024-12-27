// Package main is the entry point for the Go program.
// It demonstrates how to use conditional statements (if-else) in Go,
// including basic comparisons and logical conditions.

package main

import "fmt"

func main() {

	// Declare variables
	age := 19          // Declare an integer variable 'age' and assign the value 19
	name := "Elian"    // Declare a string variable 'name' and assign the value "Elian"

	// First conditional block: check if age is greater than or equal to 18
	if age >= 18 {
		// If the condition is true, print this message
		fmt.Println("You are old enough to drive 1")
	} 
	
	// Second conditional block: check if age is less than or equal to 10
	if age <= 10 { 
		// If the condition is true, print this message
		fmt.Println("You are not old enough to drive 2")
	} else if age == 65 && name == "elian" { 
		// Else if block: checks two conditions: if age is 65 and name is "elian"
		// Both conditions must be true to execute this block
		fmt.Println("You are now very old and can drive 3")
	} else { 
		// Else block: executed when none of the above conditions are true
		fmt.Println("You are old enough to drive 4")
	}

}
