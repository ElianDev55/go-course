package main

import (
	"encoding/json"
	"fmt"

	structs "github.com/ElianDev55/go-course/strucs/structsInterface"
	"github.com/ElianDev55/go-course/strucs/vehicles"
)

func main() {

	var product structs.Product
	product.ID = 12
	product.Name = "Laptop"
	product.Type = structs.Type{
		ID : 1,
		Code: "124ds32",
		Description: "High-end laptop with powerful processor and large storage.",
	}
	product.Price = 900.00
	fmt.Print(product)
	fmt.Println()

	product2 := structs.Product{
		ID:       13,
		Name:     "Smartphone",
		Type:    structs.Type{
		ID : 1,
		Code: "124ds32",
		Description: "High-end laptop with powerful processor and large storage.",
	},
		Price:   600.00,
	}
	fmt.Print(product2)
	fmt.Println()

	product3 := structs.Product{
		ID:       14,
		Name:     "Tablet",
	} 
	fmt.Print(product3)
	fmt.Println()

	v, err := json.Marshal(product)
	fmt.Println("value --->", string(v))
	fmt.Println("error --->", err)

	fmt.Println("product2 --->", product2.TotalPrice())
	product2.AddTag("1", "2", "3")
	fmt.Println("product3 --->", product2  )

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("VEHICLES")

	car1 := vehicles.Car{
		Time: 100,
	}
	fmt.Println(car1.Distance())

	vArray := []string{"CAR", "MOTORCYCLE", "TRUCK"}

	d	:= 400
	for _, v := range vArray {
		vh, err :=  vehicles.New(v, d)

		if err != nil {
				fmt.Println(err)
		}

		fmt.Println("vehicle --->", v)
		fmt.Println( "Vehicle Distance --->", vh.Distance())
	}

}
