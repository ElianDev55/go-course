// Package main demonstrates how to use functions, anonymous functions, and higher-order functions in Go.

package main

import "fmt"

// add is a function that takes two integers and returns their sum.
func add(a, b int) int {
  return a + b
}

// multiply is a function that takes two integers and returns their product.
func multiply(a, b int) int {
  return a * b
}	

// showResult is a higher-order function that takes a function (result) and two integers as parameters.
// It returns a formatted string that shows the result of applying the function to the integers.
func showResult(result func(int, int) int, a int, b int) string {
	return fmt.Sprintf("The result of the operation is: %d", result(a, b)) // Use Sprintf to insert the result into the string
}

func main() {

	// Anonymous function: This function subtracts b from a and returns the result.
	res := func(a int, b int) int {
		return a - b
	} 

	// Demonstrate calling showResult with different functions and parameters
	// The first example uses the anonymous subtraction function.
	fmt.Println(showResult(res, 10, 3))        // Output: The result of the operation is: 7
	fmt.Println(showResult(add, 5, 7))         // Output: The result of the operation is: 12
	fmt.Println(showResult(multiply, 6, 5))    // Output: The result of the operation is: 30
}
