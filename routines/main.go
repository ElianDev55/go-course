package main

import (
	"fmt"
	"time"
)


func main () {

	go myProcess("A")
	go myProcess("B")

	time.Sleep(3 * time.Second)

	myChannel := make(chan string)
	go myProcessChannel("C", myChannel)

	result := <-myChannel
	fmt.Println(result)
	close(myChannel)



}

func myProcessChannel(p string, c chan string){
	i := 0
	for i < 25 {
		time.Sleep(150 * time.Millisecond)
		i ++
		fmt.Println("Process ",p, "Iteraction N ", i)
	}
	c <- "OK"
}


func myProcess(p string){

	i := 0

	for i < 25 {
		time.Sleep(150 * time.Millisecond)
		i ++

		fmt.Println("Process ",p, "Iteraction N ", i)

	}

}
