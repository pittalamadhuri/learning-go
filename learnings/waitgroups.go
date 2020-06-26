package learnings

import (
	"fmt"
	"sync"
	"time"
)

func groupMemer(i int, wg *sync.WaitGroup) {
	fmt.Println("inside go routine ", i)
	time.Sleep(2* time.Second)
	fmt.Println("Done with routine ", i)
	wg.Done()
}

func LearnWaitgroups() {
	noOfRoutines := 3
	fmt.Println("Going to start ",noOfRoutines)
	var wg sync.WaitGroup
	for i := 1; i<= noOfRoutines; i++ {
		wg.Add(1)
		go groupMemer(i, &wg)
	}
	wg.Wait()
	fmt.Println("All routines finished executing")
}