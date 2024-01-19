package main

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )

// func calculate(){
// 	for i := 0; i < 3; i++ {
// 		fmt.Println("iteration in func:", i)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {
// 	runtime.GOMAXPROCS(4)
// 	go calculate()
// 	for i := 0; i < 3; i++ {
// 		fmt.Println("iteration out func:",i)
// 		time.Sleep(time.Second)
// 	}
// }