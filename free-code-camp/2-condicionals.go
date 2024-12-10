package main

import "fmt"

func main() {

	age := 19
	name := "Elian"

	if age >= 18 {
    fmt.Println("You are old enough to drive 1")
  } 
	// second code block
	if (age <= 10) { // you can use parentheses or no
		fmt.Println("You are not old enough to drive 2")
  } else if age == 65 && name  ==  "elian" { // else if  have to stay in this position
		fmt.Println("You are now very old and can drive 3")
	} else { // else  have to stay in this position
    fmt.Println("You are old enough to drive 4")
  }

}
