package main

import "fmt"

func main() {
	Constants()
	BasicTypes()
	Array()
}

func Constants()  {
	const name string = "freeman"
	noname := "ke"
	fmt.Println(name, noname)
}

func BasicTypes()  {
	var number int =10
	var str string = "Hello World"
	var isGood bool = true
	var floatValue float32 = 20.0
	var doubleValue float64 = 100.0
	var uint32Value uint = 10
	var uint64Value uint64 = 100
	var intPointer *int= &number
	fmt.Println(number, str, isGood, floatValue, doubleValue, uint32Value, uint64Value, intPointer)
}

func Array() {
	var names [2] string
	var ages = [2]int{1, 2}
	var items = [...]string{"1", "2"}

	var name = names[1]

	fmt.Println(names, ages, items, name)
}