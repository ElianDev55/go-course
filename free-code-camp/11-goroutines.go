package main

import (
	"fmt"
	"sync"
	"time"
)

func say (text string, wg *sync.WaitGroup) string {
	defer wg.Done()
	return fmt.Sprintf("Hello, %s!", text)
}

func main() {

	var wg sync.WaitGroup

	fmt.Println("Hello, world!")
	wg.Add(1)
	go  fmt.Println(say("Hi there", &wg))
	
	wg.Wait()

	go func (text string) {  
		fmt.Println(text)
	}("Bye")


	time.Sleep(time.Second * 1)

}
