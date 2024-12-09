package main

import (
	"fmt"
	"strings"
)

func main() {

	var sum int64 = 0
	for i:=0; i<10; i++ {

		if i%2!=0 { // we can use the modulus operator (%) to check if a number is odd or even. If it's not, we skip it.
		sum += int64(i)
    }
		//fmt.Println(sum)
	}

	/*
		myMapStrings := map[string]string{
		"name": "Elian",
		"age":  "19",
	}

	// k = key of the map
	// v = value of the map
  // range keyword is used to iterate over the key-value pairs in a map

	for k , v := range myMapStrings { // for each like javascript :D
		fmt.Println("This the key -->", k, "This is the value -->", v)
	}
		*/

	var fruit string = ""

	for{ // this is for creat a while loop
		if fruit == "orange"{
			fmt.Println("Exit")
			break
		}
		fmt.Print("Enter a fruit: ") // this is for ask the user for input
		fmt.Scan(&fruit)
		fruit = strings.ToLower(fruit) // convert the input to lower case (to match the map keys)
	}

}
