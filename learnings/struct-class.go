package learnings

import "fmt"

type Student struct {
	Name             string
	Age              int
	TotalSubjects    int
	FinishedSubjects int
}

func (s *Student) RemainingSubjects() {
	fmt.Printf("\n %s has %d subjects remaining ", s.Name, s.TotalSubjects-s.FinishedSubjects)
}
