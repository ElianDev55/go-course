// Package main demonstrates the use of interfaces and methods in Go.
package main

import "fmt"

// figure2d is an interface that defines a method for calculating the area of a 2D figure.
type figure2d interface {
	area() float64 // Method that returns the area of the figure as a float64.
}

// square represents a square with a base (side length).
type square struct {
	base float64 // The base represents the side length of the square.
}

// rectangle represents a rectangle with a base and a height.
type rectangle struct {
	base   float64 // The base represents the width of the rectangle.
	height float64 // The height represents the height of the rectangle.
}

// area method for the square struct calculates the area of the square (base^2).
func (s square) area() float64 {
	return s.base * s.base
}

// area method for the rectangle struct calculates the area of the rectangle (base * height).
func (r rectangle) area() float64 {
	return r.base * r.height
}

// calculate accepts any figure that implements the figure2d interface and prints its area.
func calculate(f figure2d) {
	fmt.Printf("The area of the figure is: %.2f\n", f.area()) // Calls the area method for the figure passed.
}

func main() {

	// Create a square with a base of 2 units.
	mySquare := square{base: 2}
	// Create a rectangle with a base of 2 units and height of 4 units.
	myRectangle := rectangle{base: 2, height: 4}

	// Calculate and print the area of the square.
	calculate(mySquare)
	// Calculate and print the area of the rectangle.
	calculate(myRectangle)

	// Create a list of different types of data using the empty interface.
	myInterfaces := []interface{}{1, 1.4, "hola"}
	// Print the list of interfaces (a mix of integer, float, and string).
	fmt.Println(myInterfaces)
}
