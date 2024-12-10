package main

import "fmt"


type figure2d interface {
	area() float64  // a function that returns a float64

}


type square struct {
	base float64

}

type rectangle struct {
	base float64
	height float64
}


func (s square) area () float64 {
	return s.base * s.base
}

func (r rectangle) area () float64 {
  return r.base * r.height
}

func calculate(f figure2d ){
	fmt.Printf("The area of the figure is: %.2f\n", f.area())  // we call the area function of the passed figure2d type.  The area function is defined in the interface figure2d.
}

func main() {

	mySquare := square{base:2}
	myRectangle := rectangle{base: 2 , height: 4}
	calculate(mySquare)
	calculate(myRectangle)

	// list interfaces

	// You cam create a list with different types
	myInterfaces := []interface{}{1,1.4,"hola"}
	fmt.Println(myInterfaces)

}
