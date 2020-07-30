package learnings

import (
	"fmt"
	"sync"
)

var numberArray []int

func addNumberToArray(number int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	numberArray = append(numberArray, number)
	mutex.Unlock()
	wg.Done()
}
func LearnMutex() {
	wg := sync.WaitGroup{}
	var mut = sync.Mutex{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go addNumberToArray(i, &wg, &mut)
	}
	wg.Wait()
	fmt.Println("numberArray after adding numbers ", numberArray)
}
