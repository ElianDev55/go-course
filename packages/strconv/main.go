package main

import (
	"fmt"
	"strconv"
)



func main() {
	s := strconv.Itoa(-100)
	fmt.Println(s)

	s = strconv.FormatInt(100, 8)
	fmt.Println("FormatInt:", s)

	s = strconv.FormatFloat(3.14, 'f', 2, 64)
	fmt.Println("FormatFloat:", s)

	s = strconv.FormatBool(true)
	fmt.Println("FormatBool:", s)

	s = strconv.FormatUint(100, 16)
	fmt.Println("FormatUint:", s)


}
