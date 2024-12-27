// Package main demonstrates the use of arrays and slices in Go.
// It highlights the differences between arrays (fixed size) and slices (dynamic size),
// along with some common operations like appending, slicing, and iterating.

package main

import (
	"fmt"
)

func main() {

	// !! ARRAYS ARE NOT MUTABLE, BUT SLICE IS MUTABLE LOL !!
	// Arrays in Go have a fixed size, and the size is part of their type.

	// Declaring an array of fixed size 4
	var list [4]int

	// Print the length (number of elements) of the array
	fmt.Println("This is the length of the list", len(list))
	// Print the capacity (max number of elements the array can hold) of the array
	fmt.Println("This is the max for add ", cap(list))
	// Print the array
	fmt.Println(list)

	// Slices are more flexible than arrays and have a dynamic size.
	// Declaring a slice with initial values
	slice := []int{0, 1, 2, 3, 4}
	// Print the length (number of elements) of the slice
	fmt.Println("This is the length of the slice", len(slice))
	// Print the capacity (max number of elements the slice can hold)
	fmt.Println("This is the max for add ", cap(slice))
	// Print the slice
	fmt.Println(slice)

	// Accessing and printing elements and ranges of the slice
	fmt.Println(slice[0])    // Print the first element of the slice
	fmt.Println(slice[1:])    // Print the slice starting from the second element to the end
	fmt.Println(slice[1:3])   // Print the slice from the second to the third element (excluding index 3)
	fmt.Println(slice[:3])    // Print the slice from the start to the third element (excluding index 3)

	// Appending an element to the slice
	slice = append(slice, 5)
	fmt.Println(slice)        // Print the updated slice

	// Appending another slice to the original slice
	newSlice := []int{3, 4, 5}
	slice = append(slice, newSlice...) // The '...' unpacks the new slice into the append function
	fmt.Println(slice)        // Print the updated slice

	// Iterating over the slice using a for loop
	// The range keyword returns the index and value for each element
	for index, value := range slice {
		// Print the index and value of each element in the slice
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

}
