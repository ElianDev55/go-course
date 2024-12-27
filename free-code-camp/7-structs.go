// Package main demonstrates how to use structs in Go.
package main

import "fmt"

// car is a struct that represents a car with a brand and a year.
type car struct {
	brand string  // brand is the name of the car's brand
	year  int64   // year is the manufacturing year of the car
}

func main() {
	// Create a new car instance with the brand "Testla" and year 2020.
	myCar := car{brand: "Testla", year: 2020}
	fmt.Println(myCar) // Output: {Testla 2020}

	// Another way to create a car instance using a zero-value initialization.
	var otherCar car
	otherCar.brand = "Ferrari"  // Set the brand of the car
	fmt.Println(otherCar)       // Output: {Ferrari 0} (Year is zero since it's not initialized)
}
