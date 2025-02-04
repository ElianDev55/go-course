package main

import (
	"fmt"
)

func main() {
	
	var value int = 1

	slice := make([]int, 0)
	for value != 0 {
		fmt.Println("Enter a number: ")
		fmt.Scanln(&value)
		if value != 0 {
			slice = append(slice, value)
		}
	}

	for i := range slice {
		fmt.Println()
		fmt.Println(slice[i])
	}

}
