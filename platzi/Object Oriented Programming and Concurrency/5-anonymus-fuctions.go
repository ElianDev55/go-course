package main

import (
	"fmt"
	"time"
)

func main() {

	x := 4
	y := func () int {
		return x * 2
	}

	c := make(chan int)

	go func() {
		fmt.Printf("starting")
		time.Sleep(5 * time.Second)
		c <- 1
	}()

	<- c
	fmt.Println(y()) // Outputs: 8

}
