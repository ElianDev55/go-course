package main

import (
	"fmt"
	"sync"
	"time"
)

// say function accepts a text string and a pointer to WaitGroup.
// It formats the string as "Hello, {text}!" and signals that it is done.
func say(text string, wg *sync.WaitGroup) string {
	defer wg.Done() // Decrease the WaitGroup counter when the function finishes
	return fmt.Sprintf("Hello, %s!", text) // Return formatted greeting
}

func main() {

	var wg sync.WaitGroup // Create a WaitGroup to manage goroutines

	// Print initial message
	fmt.Println("Hello, world!")

	wg.Add(1) // Increment the WaitGroup counter to wait for 1 goroutine
	go fmt.Println(say("Hi there", &wg)) // Launch the goroutine to print the greeting

	wg.Wait() // Wait for the goroutine to complete

	// Launch another goroutine with an anonymous function to print a message
	go func(text string) {
		fmt.Println(text)
	}("Bye")

	// Sleep for a second to ensure the last goroutine has time to execute
	time.Sleep(time.Second * 1)
}
