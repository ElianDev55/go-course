package main

import (
	"fmt"
)


func main () {
	
	yearsOld := 20

	fmt.Println()
	fmt.Println("Operators")
	fmt.Println("----------------")
/* 	fmt.Println(yearsOld > 18)
	fmt.Println(yearsOld < 18)
	fmt.Println(yearsOld == 18)
	fmt.Println(yearsOld != 18)
	fmt.Println(yearsOld >= 18)
	fmt.Println(yearsOld <= 18)

	fmt.Println("----------------")
	fmt.Println(yearsOld < 22 && yearsOld > 2)
	fmt.Println(yearsOld < 1 || yearsOld > 30)
	fmt.Println() */

	if yearsOld > 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a child")
	}

	if value := true; value == true {
		fmt.Println("The value is true")
	}

	if yearsOld == 20 {
		fmt.Println("You are 20 years old")
	} else if yearsOld == 21 {
		fmt.Println("You are not 20 years old")
	}

	switch yearsOld {
	case 20:
		fmt.Println("You are 20 years old")
	case 21:
		fmt.Println("You are 21 years old")
	default:
		fmt.Println("You are not 20 or 21 years old")
	}

}
