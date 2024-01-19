package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func myGoroutine(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	time.Sleep(time.Second)
// 	fmt.Println("Goroutine", id, "finish")
// }

// func main() {
// 	var wg sync.WaitGroup

// 	for i := 0; i < 3; i++ {
// 		wg.Add(1)
// 		go myGoroutine(i, &wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("all goroutine finished")
// }