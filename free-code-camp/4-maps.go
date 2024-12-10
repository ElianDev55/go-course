package main

import "fmt"

func main() {

	// Maps is similar to  dictionaries on python or object in javascript, oks?

	myMapStrings := map[string]string{
		"name": "Elian",
		"age":  "30",
	}

	fmt.Println(myMapStrings)
	fmt.Println("My name is", myMapStrings["name"] )// print the name

	myMapStrings["date"] = "2015" // add the new element
	fmt.Println(myMapStrings) 

	delete(myMapStrings, "age") // remove the old element
	fmt.Println(myMapStrings)

}
