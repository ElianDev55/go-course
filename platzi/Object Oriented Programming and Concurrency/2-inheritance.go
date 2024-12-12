package main

import "fmt"

type Person struct {
  name string
  age  int
}

type Employee struct {
  id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "John Doe"
	ftEmployee.age = 30
	ftEmployee.id = 123456
	fmt.Printf("%v", ftEmployee)
}
