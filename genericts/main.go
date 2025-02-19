package main

import "fmt"

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Printf("Sum of integers: %v\n", Sum(numbers))


    decimals := []float64{1.5, 2.5, 3.5}
    fmt.Printf("Sum of float64: %v\n", Sum2(decimals))

    
    float32Numbers := []float32{1.1, 2.2, 3.3}
    fmt.Printf("Sum of float32: %v\n", Sum(float32Numbers))

		Any("Hola", "M")

		ComparableType("1", "1")

}

func Sum[N int | int32 | float32 | float64](v []N) N {
    var total N
    for _, tV := range v {
        total += tV
    }
    return total
}


type Number interface {
	int | int32 | float32 | float64
}

type CustomSlice[V int | string] []V

func Sum2[N Number](v []N) N {
    var total N
    for _, tV := range v {
        total += tV
    }
    return total
}


func Any [N any] (v1, v2 N)  {
	fmt.Println(v1,v2)
}

func ComparableType [N comparable] (v1, v2 N){
	fmt.Println(v1,v2)
	fmt.Println(v1 == v2)
}

