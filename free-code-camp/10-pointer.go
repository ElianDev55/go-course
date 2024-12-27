// Package main demonstrates the usage of structs, methods, and pointers in Go.
package main

import (
	"fmt"
)

// pc represents a computer with RAM, disk space, and a brand.
type pc struct {
	ram   int    // Amount of RAM in GB.
	disk  int    // Disk size in GB.
	brand string // Brand of the PC.
}

// ping is a method associated with the pc struct. It prints a message and the brand of the PC.
func (myPc pc) ping() {
	fmt.Println("Pinging...")
	fmt.Println("name is", myPc.brand) // Prints the brand of the PC.
}

// duplicateRam is a method that modifies the RAM of the PC by doubling its value.
// It operates on a pointer to the pc struct, so it changes the original struct.
func (myPc *pc) duplicateRam() {
	myPc.ram = myPc.ram * 2 // Doubles the RAM value.
}

func main() {

	// Example of working with pointers in Go.
	a := 10
	b := &a // b is a pointer to a.

	// Print the value of a, its address, and the value through the pointer.
	fmt.Println("Value of a is:", a)
	fmt.Println("Address of a is:", b)
	fmt.Println("Value of a through the pointer is:", *b)

	// Change the value of a through the pointer b.
	*b = 20
	fmt.Println("Value of a changed through the pointer is:", a)

	// Create an instance of the pc struct with initial values.
	myPC := pc{ram: 8, disk: 1000, brand: "ASUS"}

	// Print the initial values of the PC.
	fmt.Println("My PC:", myPC)

	// Call the ping method to display the brand of the PC.
	myPC.ping()

	// Print the PC details after calling the ping method.
	fmt.Println("My PC:", myPC)

	// Call the duplicateRam method to double the RAM.
	myPC.duplicateRam()

	// Call the ping method again to show the updated details.
	myPC.ping()

	// Print the final details of the PC.
	fmt.Println("My PC:", myPC)
}
