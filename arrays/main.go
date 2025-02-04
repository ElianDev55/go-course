package main

import (
	"fmt"
	"unsafe"
)


func main() {

	// VECTORS
	var myintVar int 
	fmt.Println("--------> ", myintVar)
	fmt.Printf("type: %T, bytes: %d, bites: %d\n", myintVar, unsafe.Sizeof(myintVar), unsafe.Sizeof(myintVar)*8)

	var myArray [5]int
	fmt.Println("--------> ", myArray)

	myArray1 := [5]int{1, 4, 100, 42, 66}
	fmt.Println("--------> ", myArray1)

	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	myArray[3] = 4
	myArray[4] = 5
	fmt.Println("--------> ", myArray)

	fmt.Println("Size of myArray: ", len(myArray))


	// SLICES

	var slice1 [] int 
	fmt.Println("--------> ", slice1)
	fmt.Printf("type: %T, bytes: %d, bites: %d\n", slice1, unsafe.Sizeof(slice1), unsafe.Sizeof(slice1)*8)

	slice1 = append(slice1, 1, 2, 3 ,4 ,5 ,6)
	fmt.Println("--------> ", slice1)

	mySlice := myArray1[1:3] // run slice from 1 to 3
	fmt.Println("--------> ", mySlice)
	fmt.Println("Size of mySlice: ", len(mySlice))
	fmt.Println(&myArray1[2])
	fmt.Println(&mySlice[1])

	mySlice[0] = 100000
	fmt.Println("--------> ", myArray1)
	fmt.Println("--------> ", mySlice)

	fmt.Println(myArray1[:4])
	fmt.Println(myArray1[2:])

	slice := make([]int, 5,)
	fmt.Println("--------> ", slice)


}
