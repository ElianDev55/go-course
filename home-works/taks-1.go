package main

import (
	"fmt"
)

func main () {

	license := true
	age := 20

	if age > 15 && license {
		fmt.Println("You can drive")
	} else if age < 15 && !license {
		fmt.Println("You can't drive")
	}

}
