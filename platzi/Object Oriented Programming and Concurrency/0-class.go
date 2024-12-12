package main

import "fmt"

type Employee struct {
	id int 
	name string
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string)  {
  e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
  return e.name
}

func main() {

	/* 
	e := Employee {
		id: 1,
    name: "John Doe",
	}
		*/

		e := Employee{}
		e.id = 1
		e.name = "John Doe"
		fmt.Println(e) // Output: {1 John Doe}
		e.SetId(2)
		e.SetName("Jane Smith")
		fmt.Println(e) // Output: {2 Jane Smith}
		fmt.Println(e.GetId()) // Output: 2
		fmt.Println(e.GetName()) // Output: Jane Smith
}

