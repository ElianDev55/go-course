package main

import (
	"fmt"

	"github.com/ElianDev55/go-course/home-works/taks-4/matrix"
)



func main () {


	fmt.Println("hello world")

		m, err := matrix.New([]int64{2, 7, 8, 5}, []int64{4, 4, 7}, []int64{5, 6, 1})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	m.Print()

}
