package fuction

import (
	"errors"
	"fmt"
)

type Operation int

const (
	SUM Operation =  iota
	SUB
	DIV
	MUL

)

func Add(a int, b int) int {
	return a + b
}

func RepeatString(incremet int , value string) {
	for i := 0; i < incremet; i++ {
		fmt.Println(value)
	}
}

func Calc(op  Operation, x , y float64) (float64, error) {
	switch op {
	case SUM:
		return x + y, nil
	case SUB:
		return x - y, nil
	case DIV:
		if y == 0 {
			return 0, errors.New("Can't divide by zero")
		}
		return x / y, nil
	case MUL:
		return x * y, nil
	default:
		return 0, errors.New("Invalid operation")
	}
}

func Split(v int) (int , int) {
	x := v * 4 / 9
	y := v - x
	return x, y
}

func Msum(values ...float64) float64 {
  var sum float64
  for _, v := range values {
    sum += v
  }
  return sum
}
