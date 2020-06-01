package test

import "fmt"

type employee struct {
	name string
	age int
}

func (e employee) displayDetails() {
	fmt.Println("Name : ", e.name, " age: ", e.age)
}

func (e *employee) changeAge() {
	e.age++
}

func LearnPointers() {
	b := 10
	a := &b
	emp := &employee{
		name: "Madhuri",
		age: 23,
	}

	empStruct := employee{
		name: "Madhuri",
		age: 23,
	}

	addressTemp := new(int)
	fmt.Printf("Type of pointer a is  %T ", a)
	fmt.Println("Address of b is ", a)
	fmt.Println("address created using new is ", addressTemp)
	fmt.Println("addresstemp points to ", *addressTemp)
	fmt.Println("b before incremented ",b)
	*a++
	fmt.Println("b after increment using its pointer ", b)
	fmt.Println("printing struct with pointer ", emp.name, emp.age)
	emp.displayDetails()
	emp.changeAge()
	(&empStruct).displayDetails()
	(&empStruct).changeAge()
	fmt.Println("emp age after changing is ", emp.age)
	fmt.Println("empStruct age after changing is ", empStruct.age)
}
