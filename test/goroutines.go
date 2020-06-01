package test

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i:=1; i<=5; i++ {
		fmt.Printf("\n %d",i)
		time.Sleep(1*time.Second)
	}
}

func printAlphabets() {
	for i:='a'; i<='e'; i++ {
		fmt.Printf("\n %c", i)
		time.Sleep(1*time.Second)
	}
}
func LearnGoroutines(){
	fmt.Println("started Main function")
	//time.Sleep(2* time.Second)
	go printNumbers()
	go printAlphabets()
	time.Sleep(5* time.Second)
	fmt.Println("Ending Main function")
}
