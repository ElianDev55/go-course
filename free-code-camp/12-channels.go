package main

import "fmt"

// c  <- c exit of values
//  text =  <- c

func say (text string, c chan <- string) {
	 c <- text // revice vlue
} 

func main() {

	c := make(chan string, 1) // number unit values have

	fmt.Println("Hello")

	go say("World", c)

	fmt.Println(<- c)


}
