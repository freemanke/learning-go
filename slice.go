package main

import "fmt"

func main() {
	ages := []int{1,2,4,3,4,5,6,4,6,4,3}
	slice := ages[:]
	sliceStart := ages[4:]
	sliceEnd := ages[:3]

	fmt.Println(slice)
	fmt.Println(sliceStart)
	fmt.Println(sliceEnd)

}