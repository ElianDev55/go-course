package main

import (
	"errors"
	"fmt"

	"github.com/ElianDev55/go-course/errors/operator"
)

func main () {

  var err error

  err = errors.New("This is an error")
  fmt.Println(err)
  fmt.Println(err.Error())

  err2 := fmt.Errorf("My format err, string %s, number %d", "hello", 10)
  fmt.Println(err2)

  defer func(){
    fmt.Println("Instan function run in tha final running")
    r := recover()
    if r != nil {
      fmt.Println("Error en el r")
    }
  }()

  x := 4
  y := 0
  z := operator.Div(x,y)

  fmt.Println(z)

  

}
