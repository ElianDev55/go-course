package main

import (
	"fmt"
)

func main() {

	// !! ARRAYS IS NO MUTABLE BU SLICE IS MUTABLE LOL !!

	// Arrays are fixed size, and their length is part of their type.

	var list [4]int

	fmt.Println("This is the length of the list", len(list))
	fmt.Println("This is the max for add ", cap(list))
	fmt.Println(list)

	// Slice is fixed size, and its length is part

	slice := []int{0, 1, 2, 3, 4}
	fmt.Println("This is the length of the slice", len(slice))
	fmt.Println("This is the max for add ", cap(slice))
	fmt.Println(slice)

	// Methods on the slice
	fmt.Println(slice[0])
	fmt.Println(slice[1:])
	fmt.Println(slice[1:3])
	fmt.Println(slice[:3])

	// Append to slice
	slice = append(slice, 5)
  fmt.Println(slice)

  // Append new slice
	newSlice := []int{3, 4, 5}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	// run the slice

	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

}
