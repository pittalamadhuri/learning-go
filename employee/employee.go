package employee

import "fmt"

type Employee struct {
	Name        string
	Age         int
	TotalLeaves int
	LeavesTaken int
}

func New(Name string, Age int, TotalLeaves int, LeavesTaken int) Employee {
	return Employee{
		Name,
		Age,
		TotalLeaves,
		LeavesTaken,
	}
}

func (e *Employee) RemainingLeaves() {
	fmt.Printf("\n %s has %d leaves remaining ", e.Name, e.TotalLeaves-e.LeavesTaken)
}
