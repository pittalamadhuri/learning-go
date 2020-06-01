package learnings

import "fmt"

func AddAgeWeight(age int, weight int) (tot int) {
	tot = age + weight
	return
}

func init() {
	fmt.Println("learnings is inititalized")
}
