package main

import "fmt"


func sum(values ...int) int {
	total := 0
	for _, value := range values {
    total += value
  }
	return total
}

func getValues(x int) (int, int , int)  {
	return 2 * x, 2 * x, 2 * x
}

func getValue(x int) (single int , doblue int , div int ) {
	single = x
	doblue = 2 * x
	div = x / 2
	return
}

func main() {
	fmt.Print(sum(11))
	fmt.Print("\n")
	fmt.Print(sum(1, 2, 3, 4, 5))
	fmt.Print("\n")
	fmt.Print(getValues(2))
	fmt.Print("\n")
	singl1e, doblue, div := getValue(3)
	fmt.Print(singl1e, doblue, div)
}
