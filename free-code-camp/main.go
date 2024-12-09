package main

import (
	"fmt"
	CarSt "free-code-camp/structs" // make alias for struct
)

func main() {
	var myCar CarSt.CarPublic
	myCar.Brand = "Testla"
	fmt.Println(myCar)
	CarSt.HelloWorld()
}
