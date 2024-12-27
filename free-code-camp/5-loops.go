// Package main demonstrates various Go control structures including loops, conditionals, and string manipulation.

package main

import (
	"fmt"
	"strings"
)

func main() {

	// Variable to store the sum of odd numbers.
	var sum int64 = 0

	// Loop from 0 to 9 to find the sum of odd numbers.
	// The modulus operator (%) is used to check if a number is odd or even.
	for i := 0; i < 10; i++ {
		// If the number is odd (i % 2 != 0), add it to the sum
		if i%2 != 0 {
			sum += int64(i) // Convert the integer to int64 before adding to sum
		}
		//fmt.Println(sum) // Uncomment this line if you want to see the sum at each step
	}

	// A commented-out code block that demonstrates how to use a map and iterate over its key-value pairs
	/*
		myMapStrings := map[string]string{
			"name": "Elian",
			"age":  "19",
		}

		// k = key of the map
		// v = value of the map
		// range keyword is used to iterate over the key-value pairs in a map
		for k, v := range myMapStrings { // Iterate over the map
			fmt.Println("This is the key -->", k, "This is the value -->", v)
		}
	*/

	// Variable to store the fruit entered by the user.
	var fruit string = ""

	// Infinite loop, acts as a "while" loop in Go
	// It will continue until the user enters "orange"
	for {
		if fruit == "orange" { // If the fruit entered is "orange", exit the loop
			fmt.Println("Exit")
			break // Break out of the loop
		}
		// Prompt the user to enter a fruit
		fmt.Print("Enter a fruit: ")
		// Get user input for the fruit
		fmt.Scan(&fruit)
		// Convert the fruit input to lowercase to match the map keys (case-insensitive)
		fruit = strings.ToLower(fruit)
	}

}
