package main

import (
	"fmt"
)

func main() {

	array := []int{4,2,5,6,7}

	for i := range array {
		value := array[i] * 2
		array[i] = value
	}

	fmt.Println(array)

}
