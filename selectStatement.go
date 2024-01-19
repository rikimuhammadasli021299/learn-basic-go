package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 42
	}()
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- 100
	}()

	select{
	case value := <-ch1:
		fmt.Println("Receive from ch1", value)
	case value := <-ch2:
		fmt.Println("Receive from ch2", value)
	}
}