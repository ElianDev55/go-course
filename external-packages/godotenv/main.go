package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)



func main() {
	s := os.Getenv("EXAMPLE_STRING")
	fmt.Println(s)

	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		return
	}

	s = os.Getenv("EXAMPLE_STRING")
	fmt.Println(s)

	
	/* if err := godotenv.Load("uuid/.var"); err != nil {
		fmt.Printf("Error loading uuid/.var file: %v\n", err)
		return
	}

	in := os.Getenv("EXAMPLE_INT")
	fmt.Println(in)
	 */

	myEnv, err := godotenv.Read()
	if err != nil {
		fmt.Printf("Error reading .env file: %v\n", err)
	}

	fmt.Println(myEnv)

}
