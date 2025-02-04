package main

import (
	"fmt"
)

func  main() {

	sum := 0

	for i := 0; i < 10; i++ {
		sum ++
	}

	fmt.Println(sum)

	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)

	for {
		sum += sum
		if sum > 1000 {
			break
		}
	}

	arr := []int{1, 2, 3, 4, 5}
	
	fmt.Println()

	for i := range arr {
		fmt.Println("index:", i , "value:", arr[i])
	}

	for i, v := range arr {
		fmt.Println("index:", i , "value:", v)
	}

	fmt.Println("Run for loop with map")
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	for k, v := range m {
		fmt.Println("key:", k, "value:", v)
	}

	mapWithSlice := map[string][]int{"a": {1, 2, 3}, "b": {4, 5, 6}, "c": {7, 8, 9}}

	for k, v := range mapWithSlice {
		fmt.Println("key:", k, "value:", v)
		for i, val := range v {
			fmt.Println("index:", i, "value:", val)
		}
	}

}
