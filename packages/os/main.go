package main

import (
	"fmt"
	"os"
)



func main() {

	file , err := os.Open("file.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := make([]byte, 100)

	c, err := file.Read(data)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%d bytes read: %s\n", c, string(data[:c]))
	file.Close()

	USERNAME := os.Getenv("USERNAME")
	fmt.Println(USERNAME)

	os.Setenv("PASSWORD", "password")
	PASSWORD := os.Getenv("PASSWORD")
	fmt.Println(PASSWORD)

	dir, _ := os.Getwd()
	fmt.Println(dir)
	
	os.Setenv("EMPTY", "")
	EMPTY := os.Getenv("EMPTY")
	fmt.Println(EMPTY)

	LOL := os.Getenv("LOL")
	fmt.Println(LOL)

	env, ok := os.LookupEnv("EMPTY")
	fmt.Println(env, ok)

	env1, ok1 := os.LookupEnv("LOL")
	fmt.Println(env1, ok1)

	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "password")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")

	// ${value} or $value
	dbUrl := os.ExpandEnv("$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/mydb")
	fmt.Println(dbUrl)



}
