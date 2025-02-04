package main

import "fmt"




func main () {

	var i int 
	var iP *int 

	i = 34
	iP = &i

	fmt.Println("Value --->", i)
	fmt.Println("Address--->" , &i)
	fmt.Println("Pointer without * --->", iP)
	fmt.Println("Pointer with * --->", *iP)

	*iP = 56
	fmt.Println("New Value --->", i) 
	fmt.Println("New Address--->" , &i) 
	fmt.Println("New Pointer without * --->", iP)
	fmt.Println("New Pointer with * --->", *iP)

}
