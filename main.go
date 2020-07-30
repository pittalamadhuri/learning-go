package main

import (
	"fmt"
	"learngo/employee"
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
	learnings.LearnSelect()
	learnings.LearnMutex()
	// Go is not purely Object oriented, some of the identical natures of Go with Object orientation are below

	//Using structs for classes
	structs_classes()

	//Using Composition for inheritance
	composition_inheritance()

	//Using structs for polymorphism
	structs_polymorphism()

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

func structs_classes() {
	student1 := learnings.Student{
		Name:             "Madhuri",
		Age:              24,
		FinishedSubjects: 2,
		TotalSubjects:    10,
	}

	student1.RemainingSubjects()

	student2 := learnings.Student{} //Object creation is possible without default values in go, hence we need a constructor for the object to be valid

	student2.RemainingSubjects()

	//Hence we create a New() function for initializing the object to avoid invalid/zero objects.. this is done in new package(just for learning)
	employee1 := employee.New("Tom", 25, 10, 4)
	employee1.RemainingLeaves()
}

type pencil struct {
	brand string
	cost  string
}

func (p *pencil) Writing() {
	fmt.Printf("\n %s pencil is writing", p.brand)
}

// ipad here is composed of ipad properties and also the pencil properties, it can be similar to inheriting a pencil class
type ipad struct {
	gen  string
	cost string
	pencil
}

func (i *ipad) Existing() {
	fmt.Printf("\n %s gen ipad is Existing with %s pencil", i.gen, i.brand)
}

//slices inside struct is also possible

type page struct {
	text   string
	pageNo int
}

type book struct {
	title  string
	author string
	pages  []page
}

func (b *book) describeBook() {
	fmt.Printf("\n %s book has %d pages ", b.title, len(b.pages))
}

func composition_inheritance() {
	ipad1 := ipad{
		gen:  "7th",
		cost: "1000",
		pencil: pencil{
			brand: "Apple",
			cost:  "100",
		},
	}
	ipad1.Writing()
	ipad1.Existing()

	page1 := page{
		text:   "How I met the kite Runner",
		pageNo: 1,
	}
	page2 := page{
		text:   "Running kites",
		pageNo: 13,
	}
	var book1 = book{
		title:  "Kite Runner",
		author: "Khaled",
		pages:  []page{page1, page2},
	}
	book1.describeBook()
}

type fixedTime struct {
	project      string
	assignedTime int
}

func (f *fixedTime) timeTaken() int {
	return f.assignedTime
}

type variableTime struct {
	project     string
	files       int
	timePerFile int
}

func (v *variableTime) timeTaken() int {
	return v.files * v.timePerFile
}

func structs_polymorphism() {
	//there are 2 different structs but have the same method implemented based on the time, hence it is identical to polymorphism

	time1 := fixedTime{
		project:      "Development",
		assignedTime: 8,
	}
	time2 := variableTime{
		project:     "Presales",
		files:       12,
		timePerFile: 1,
	}

	totalTime := time1.timeTaken() + time2.timeTaken()
	fmt.Println("\n total time taken is ", totalTime)
}
