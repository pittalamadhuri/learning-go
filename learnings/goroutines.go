package learnings

import (
	"fmt"
	"time"
)

func printNumbers(channel2 chan string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("\n %d", i)
		time.Sleep(1 * time.Second)
	}
	channel2 <- "integers done"
}

func printAlphabets(channel1 chan int) {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("\n %c", i)
		time.Sleep(1 * time.Second)
	}
	channel1 <- 1
}

func loopNumbers(numberChannel chan int) {
	for i := 0; i<=5 ; i++ {
		numberChannel <- i
	}
}

func LearnGoroutines() {
	fmt.Println("started Main function")
	//time.Sleep(2* time.Second)
	//channel declaration
	var channel1 chan int
	if channel1 == nil {
		fmt.Printf("channel1 is nil, hence defining it but the type is already %T\n", channel1)
		//channel definition
		channel1 = make( chan int)
		fmt.Printf("channel1 is now of type %T\n", channel1)
		fmt.Println("channel value now is ", channel1)
	}
	channel2 := make(chan string)
	fmt.Printf("now declared channel channel2 with type %T\n", channel2)
	go printNumbers(channel2)
	go printAlphabets(channel1)
	fmt.Println("Reading data from channel")
	dataIntergers:= <- channel2
	dataAlphabets:= <- channel1
	// the below sleep is not needed now to wait for goroutines to execute, as it waits till someone writes to the channel already YAY!!
	//time.Sleep(5 * time.Second)

	//loop channel data, receiving
	numberChannel := make(chan int)
	go loopNumbers(numberChannel)
	fmt.Println("called loop numbers")
	//using for loop
	for i := 0; i<=5 ; i++ {
		 fmt.Println("for loop coming from channel ", <- numberChannel)
	}


	//using infinite loop which breaks when there is no data any more in the channel
	//for {
	//	v, ok := <- numberChannel
	//	fmt.Println("infinite loop coming from channel ", v, "ok? ", ok)
	//	if ok == false {
	//		break
	//	}
	//}
	fmt.Println("\nEnding Main function with data in channel2 ", dataIntergers, " channel1 is ", dataAlphabets)
	learnBufferedChannels()
}

func learnBufferedChannels() {
	buffChannel := make(chan int, 3)
	buffChannel <- 1
	buffChannel <- 2
	fmt.Printf("\nbefore reading capacity of channel is %d and length is %d", cap(buffChannel), len(buffChannel))
	fmt.Println("\n1st value is ", <-buffChannel, " 2nd value is ", <-buffChannel)
	fmt.Printf("after reading capacity of channel is %d and length is %d", cap(buffChannel), len(buffChannel))
}