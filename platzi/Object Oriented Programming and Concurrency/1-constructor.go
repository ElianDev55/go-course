package main

import (
	"fmt"
)

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
    name:     name,
    vacation: vacation,
	}
}

func main() {
	e := Employee{}
	fmt.Println(e)

	e1 := Employee{
		id:       1,
		name:     "John Doe",
		vacation: true,
	}

	e2 := new (Employee)
	e2.id = 2
	e2.name = "Jane Doe"
	e2.vacation = false

	e3 := NewEmployee(
		3,
    "Alice Doe",
    false,
	)

	fmt.Println(e1)
	fmt.Println(e3)
}
