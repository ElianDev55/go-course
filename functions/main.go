package main

import (
	"fmt"

	"github.com/ElianDev55/go-course/functions/fuction"
)

func main() {
	result := fuction.Add(1, 2)
	fmt.Println(result)
	fuction.RepeatString(3, "Hello")
	value, error :=fuction.Calc(fuction.DIV, 1, 0)
	if (error != nil) {
		fmt.Println(error.Error())
	}else {
		fmt.Println(value)
	}

	fmt.Println(value, error)
	x, y := fuction.Split(20)
	fmt.Println(x,y)
	fmt.Println(fuction.Msum(1,2,3,4,5,6,7,8,97,8))
	fn := fuction.FactoryOperacion(fuction.SUM)
	fmt.Println(fn(1,2))
}
