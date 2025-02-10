package main

import (
	"fmt"
	"log"
	"os"
	"time"
)



func main(){

  date := time.Now()

	log.Println("Hello, World!") // 2025/02/05 13:54:45 Hello, World!
  //log.Fatal("This is a fatal error") // 2025/02/05 13:54:45 This is a fatal error
	//log.Panic("This is a panic error") // 2025/02/05 13:54:45 This is a panic error
  log.Printf("The value of x is %d", 42) // 2025/02/05 13:54:45 The value of x is 42
  
  file , err := os.Create(fmt.Sprintf("%d.log", date.UnixNano()))
	if err != nil {
    panic(err)
	}
  
  l := log.New(file, "Successful: ", log.LstdFlags|log.Lshortfile)
  l.Println("Logging to a file") 
  //Successful: 2025/02/05 14:01:47 main.go:22: Logging to a file

  l.SetPrefix("Error: ")
  l.Println("This is an error message")


}
