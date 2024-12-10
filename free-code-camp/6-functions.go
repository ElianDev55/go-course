package main

import "fmt"

func add(a, b int) int {
  return a + b
}

func multiply(a, b int) int {
  return a * b
}	

func showResult( result func(int,int) int, a int , b int ) string {
	return fmt.Sprintf("The result of the sum is: %d", result(a , b)) // we have to use sprintf for add numbers on the string	
}


func main() {

	// Anonymous functions
	res := func (a int , b int) int {
		return a-b
	} 

	// i have to give the fuction without variables onsite 
	fmt.Println(showResult(res, 10, 3))       // Output: The result is: 7
	fmt.Println(showResult(add, 5, 7))        // Output: The result is: 12
	fmt.Println(showResult(multiply, 6, 5))   // Output: The result is: 30


}
