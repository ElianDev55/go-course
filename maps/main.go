package main

import (
	"fmt"
)


func main() {

	map1 := make(map[string]string)

	map1["name"] = "Elian"
	map1["age"] = "20"

	fmt.Println(map1)
	fmt.Println(map1["name"])
	fmt.Println(map1["age"])

	delete(map1, "name")
	fmt.Println(map1)

	map1["name"] = "Elian"
	map1["name"] = "Joan"
	fmt.Println(map1)

	map1["empty"] = ""

	v, ok := map1["empty"]
	fmt.Println("The value:", v, "Present?", ok)

	v, ok = map1["400"]
	fmt.Println("The value:", v, "Present?", ok)

	map2 := map[string]string{
		"name": "Elian",
		"age": "20",
	}

	fmt.Println(map2)


}
