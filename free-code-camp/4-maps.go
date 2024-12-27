// Package main demonstrates how to use maps in Go.
// Maps in Go are similar to dictionaries in Python or objects in JavaScript.

package main

import "fmt"

func main() {

	// Declare and initialize a map where the key and value are both strings.
	myMapStrings := map[string]string{
		"name": "Elian",   // Key "name" with value "Elian"
		"age":  "30",      // Key "age" with value "30"
	}

	// Print the entire map
	fmt.Println(myMapStrings)

	// Access and print a specific value from the map using its key
	// Here, we print the value associated with the key "name"
	fmt.Println("My name is", myMapStrings["name"])

	// Add a new key-value pair to the map
	myMapStrings["date"] = "2015" // Add "date" with value "2015"
	// Print the updated map
	fmt.Println(myMapStrings)

	// Remove a key-value pair from the map
	delete(myMapStrings, "age") // Remove the key "age"
	// Print the updated map after deletion
	fmt.Println(myMapStrings)

}
