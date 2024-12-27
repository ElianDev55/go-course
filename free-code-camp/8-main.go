// Package main demonstrates how to use structs and import them with an alias in Go.
package main

import (
	"fmt"
	CarSt "free-code-camp/structs" // Import the package "structs" and create an alias "CarSt" for it
)

func main() {
	// Create an instance of the CarPublic struct from the imported package
	var myCar CarSt.CarPublic
	myCar.Brand = "Testla" // Set the "Brand" field of the CarPublic struct to "Testla"

	// Print the myCar struct, which includes the brand and any other fields in the struct
	fmt.Println(myCar)

	// Call the HelloWorld function from the CarSt package
	CarSt.HelloWorld()
}
