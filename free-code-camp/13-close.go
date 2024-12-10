package main

import "fmt"


func message(s string, c chan <- string) string {
	c <- s
	return s
}

func main() {

	c := make(chan string, 2)
	c <- "Hello"
	c <- "World"

	fmt.Println(len(c), cap(c))

	// Range and Close

	close(c)

	for message := range c {
		fmt.Println(message)
	}

	// Slect 

	email1 := make(chan string)
	email2 := make(chan string)

	go message("Email 1", email1)
	go message("Email 2", email2)

	for i:=0; i < 2; i++ {
		select{

		case m1 := <-email1:
			fmt.Println("Received from email1: ", m1)
		case m2 := <-email2:
			fmt.Println("Received from email2: ", m2)

		}
	}


}
