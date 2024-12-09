package main

import (
	"fmt"
)

type pc struct {
	ram int 
	disk int
	brand string
}

func (myPc pc) ping() {
	fmt.Println("Pinging...")
	fmt.Println("name is",myPc.brand)
}

func  (myPc *pc) duplicateRam()  {
	myPc.ram = myPc.ram * 2
}

func main() {

	a := 10
	b := &a

	fmt.Println("Value of a is:", a)
	fmt.Println("Address of a is:", b)
	fmt.Println("Value of a through the pointer is:", *b)

	*b = 20
	fmt.Println("Value of a changed through the pointer is:", a)

	myPC := pc{ram: 8, disk: 1000, brand: "ASUS"}
	fmt.Println("My PC:", myPC)
	myPC.ping()
	fmt.Println("My PC:", myPC)
	myPC.duplicateRam()
	myPC.ping()
	fmt.Println("My PC:", myPC)

}
