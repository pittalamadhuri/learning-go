package main

import (
	"fmt"
	"learngo/learnings"
	"reflect"
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello World")
	const gender = "female"
	var age, weight = 10, 20
	var (
		height = 50
	)
	location := "Hyderabad"
	fmt.Println("my age is ", age)
	fmt.Println("my weight is ", weight)
	fmt.Println("my height is ", height)
	fmt.Println("My location is ", location)
	fmt.Printf("My gender is %s, %T \n", gender, gender)
	fmt.Println("Infered type is ", reflect.TypeOf(age))
	tot := learnings.AddAgeWeight(age, weight)
	fmt.Println("add Age weight is ", tot)
	if class := 20; class/10 == 2 {
		fmt.Println("class is 20")
	} else {
		fmt.Println("class is not 20")
	}
	var a [5]int
	fmt.Println("a[0] is ", a[0])
	var slicea = a[1:4]
	fmt.Println("slice a is ", slicea)
	slicea[2]++
	fmt.Println("array after increment slice a is ", a)
	slice2 := []int{2, 3}
	slice2 = append(slice2, 4)
	fmt.Println("slice 2", slice2)
	myVariadic(age, weight, height)
	mapFunction()
	byteString("Hello world")
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	byteToString(byteSlice)
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	runeToString(runeSlice)
	learnings.LearnPointers()
	learnings.LearnInterfaces()
	learnings.LearnGoroutines()
	learnings.LearnWaitgroups()
	learnings.LearnWorkerPools()
}

func init() {
	fmt.Println("Main is initialized")
}

func myVariadic(age int, otherVars ...int) {
	fmt.Println("age is ", age)
	fmt.Println("variable is ", otherVars)
}

func mapFunction() {
	salaryMap := map[string]int{
		"Sowmya": 20000,
	}
	salaryMap["madhuri"] = 10000
	fmt.Println("salary of madhuri is ", salaryMap["madhuri"])
	fmt.Println("salary map contents ", salaryMap)
	value, ok := salaryMap["amith"]
	if ok != false {
		fmt.Println("salary of Amith is ", value)
	} else {
		fmt.Println("AMith salary is not provided")
	}
}

func byteString(temp string) {
	fmt.Printf("string byte array is %x \n", temp[0])
}

func byteToString(byteSlice []byte) {
	str := string(byteSlice)
	fmt.Println("byte  converted is ", byteSlice)
	fmt.Println("string converted is ", str)
}

func runeToString(runeSlice []rune) {
	str := string(runeSlice)
	fmt.Println("rune to string is ", str)
	findRuneCount(str)
	mutedString := mutateRune([]rune(str))
	fmt.Println("after mutating rune str is ", string(mutedString))
	//str[0] = 'a'
}

func findRuneCount(str string) {
	lenRune := utf8.RuneCountInString(str)
	fmt.Printf("rune count in string %s is %d \n", str, lenRune)
	fmt.Printf("char count in string %s is %d \n", str, len(str))
}

func mutateRune(runeStr []rune) []rune {
	runeStr[0] = 'a'
	return runeStr
}
