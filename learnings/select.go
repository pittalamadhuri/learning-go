package learnings

import (
	"fmt"
	"time"
)

func immitateFastServer(ch chan string) {
	time.Sleep(time.Second*2)
	ch <- "fast server done"
}

func immitateSlowServer(ch chan string) {
	time.Sleep(time.Second*5)
	ch <- "slow server done"
}

func immediateServer1(ch chan string) {
	ch <- "immediateServer1 executed"
}

func immediateServer2(ch chan string) {
	ch <- "immediateServer2 executed"
}

func LearnSelect (){
	var ch1 = make(chan string)
	var ch2 = make(chan string)
	var dummyChannel = make(chan string)
	var dummyChannel2 chan string
	go immitateFastServer(ch1)
	go immitateSlowServer(ch2)

	channel1 := make(chan string)
	channel2 := make(chan string)
	go immediateServer1(channel1)
	go immediateServer2(channel2)

	time.Sleep(time.Second * 1)

	//select when random case gets executed as both of the cases satisfy
	select {
	case chOutput1 := <- channel1 : fmt.Println(chOutput1)
	case chOutput2 := <- channel2 : fmt.Println(chOutput2)
	}

	//select with no default waits for the cases for ever
	select {
	case output := <- ch1 : fmt.Println("received from channel 1 ",output)
	case output := <-ch2 : fmt.Println("received from channel 2 ", output)
	//default is executed when at the time of select execution, both of the cases are not satisfied
	//default prevents deadlock when none of the case is satisfied
	default:
		fmt.Println("Default case executed")
	}

	//the below select never executes(when default is commented) as it is waiting to read from a channel, no go routine writes to, hence a deadlock
	select {
	 case output1 := <- dummyChannel : fmt.Println("this will never execute haha ",output1)
	default:
		fmt.Println("nothing is in dummyChannel hence default executed")
	}

	fmt.Println("value in dummyChannel2 ", dummyChannel2)
	select {
	case output3 := <- dummyChannel2 : fmt.Println("this will also never execute", output3)
	default:
		fmt.Println("dummy channel 2 is nil hence default executed")
	}

	//use select inside a for loop with a default case for a non blocking select :)
}

