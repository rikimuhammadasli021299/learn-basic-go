package main

import (
	"fmt"
	"time"
)

func print(ch chan int) {
	ch <- 20 // send value
}

func print_two(ch chan int) {
	value := <-ch // receive from go routine 1
	fmt.Println(value)
	time.Sleep(time.Second)
}

func main() {
	ch := make(chan int)
	go print(ch)
	go print_two(ch)
	time.Sleep(time.Second)
}