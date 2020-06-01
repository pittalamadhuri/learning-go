package test

import (
	"fmt"
	"reflect"
)

type employeeInterface interface {
	works() bool
}

type permEmployee struct {
	empType string
	isWorks bool
}

func (pe permEmployee) works() bool {
	return pe.isWorks
}

type tempEmployee struct {
	empType string
	isWorks bool
}

func (te tempEmployee) works() bool {
	return te.isWorks
}

func checkWorks(ei employeeInterface) {

	fmt.Println("employee works is ", ei.works())
	fmt.Printf("value inside emp is %V \n", ei)
	fmt.Printf("type of interface is %T \n", ei)
}

func getIntFromInterface(i interface{}) {
	fmt.Println("int in interface is ", i.(int))
}

func getStringFromInterface(i interface{}) {
	fmt.Println("striing in interface is ", i.(string))
	fmt.Println("Type of interface is ", reflect.TypeOf(i))
}

func LearnInterfaces() {
	emp1 := tempEmployee{
		empType: "temp employee",
		isWorks: false,
	}
	emp2 := permEmployee{
		empType: "Perm employee",
		isWorks: true,
	}
	intType := 10
	checkWorks(emp1)
	checkWorks(emp2)
	getIntFromInterface(intType)
	getStringFromInterface("this is string")
}