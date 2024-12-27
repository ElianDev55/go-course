// Package main is the entry point for the Go program.
// It demonstrates basic variable types in Go, including strings, integers, and lists (arrays and slices).
// The program also shows how to manipulate and print these types.

package main

// The "fmt" package is used for formatted I/O operations in Go.
// In this program, it's used to print messages to the console.
import "fmt"

func main() {
	// String variables
	var name string = "Elian"                // Declare a string variable 'name' and assign it the value "Elian"
	var lastName = "Villamarin"              // Declare a string variable 'lastName' with type inferred as string
	lastNameTwo := "Urrutia"                 // Declare a string variable 'lastNameTwo' using short variable declaration

	// Integer variable
	var age int8 = 19                        // Declare an 8-bit integer variable 'age' and assign the value 19
	year := 2005                             // Declare an integer 'year' using short variable declaration

	// Array with fixed size (4 elements) of strings
	var listFruts = [4]string{"Apple", "Orange"}  // Declare an array 'listFruts' with a limit of 4 elements of type string
	listCountries := [4]string{}              // Declare an empty array 'listCountries' of 4 strings
	listCountries[0] = "Colombia"            // Assign a value to the first element of 'listCountries'
	listCountries[1] = "Ecuador"             // Assign a value to the second element of 'listCountries'
	listCountries[1] = "Usa"                 // Assign a value to the second element again (overwriting "Ecuador")

	// Slice (dynamic-sized array) without a limit
	var listCountriesUnlimited []string       // Declare an empty slice of strings
	listCountriesUnlimited = append(listCountriesUnlimited, "Colombia") // Add "Colombia" to the slice
	listCountriesUnlimited = append(listCountriesUnlimited, "Ecuador") // Add "Ecuador" to the slice
	listCountriesUnlimited = append(listCountriesUnlimited, "Usa")    // Add "Usa" to the slice

	// Slices with ranges
	listCountriesTwo := listCountriesUnlimited[0:2]  // Slice the list from index 0 to index 2 (not including 2)
	listCountriesTheree := listCountriesUnlimited[1:] // Slice the list starting from index 1 to the end
	listCountriesFour := listCountriesUnlimited[:2]   // Slice the list from the beginning to index 2 (not including 2)

	// Printing the variables and lists
	fmt.Println("hello world i am " + name + " " + lastName + " " + lastNameTwo) // Print a concatenated string of names
	fmt.Println("I am ", age, " years old") // Print the integer 'age'
	fmt.Println("I was born in ", year)     // Print the integer 'year'
	fmt.Println("List of", listFruts)       // Print the entire array 'listFruts'
	fmt.Println("List of", listFruts[1])    // Print the second item of the array 'listFruts'
	fmt.Println("List of countries", listCountries) // Print the array 'listCountries'
	fmt.Println("List of countries unlimited", listCountriesUnlimited) // Print the slice 'listCountriesUnlimited'
	fmt.Println("List of countries two", listCountriesTwo) // Print the slice 'listCountriesTwo'
	fmt.Println("List of countries theree", listCountriesTheree) // Print the slice 'listCountriesTheree'
	fmt.Println("List of countries four", listCountriesFour) // Print the slice 'listCountriesFour'
}
