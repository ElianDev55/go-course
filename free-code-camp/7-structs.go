package main

import "fmt"

type car struct {
	brand string
	year  int64
}

func main() {
	myCar := car{brand: "Testla", year: 2020}
	fmt.Println(myCar) // {Testla 2020}

	// another do
	var otherCar car
	otherCar.brand= "Ferrari"
	fmt.Println()

}
