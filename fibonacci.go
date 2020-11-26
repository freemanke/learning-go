package main

import (
	"fmt"
)

func main() {
	fmt.Println("fibonacci")
	fmt.Println(fibonacci(30))
}

func fibonacci(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}