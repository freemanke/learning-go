package main

import (
	"fmt"
	"time"
)

func main() {
	go DoSomeThing(1)
	go DoSomeThing(2)
	DoSomeThing(3)
}

func DoSomeThing(number int) {
	for i:=0;i<10;i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("Goroutine #%d, Ticks: %d Date: %s\n", number, i, time.Now())
	}
}