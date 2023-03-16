package main

import (
	"fmt"
)

func main() {
	done1 := make(chan bool)
	done2 := make(chan bool)

	data1 := "bisa"
	data2 := "coba"

	go Goroutine1(data1, done1)
	go Goroutine2(data2, done2)

	<-done1
	<-done2

	fmt.Println("Both goroutines have finished")
}

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// var mutex = &sync.Mutex{}

// func goroutine1(data interface{}, done chan bool) {
// 	for i := 0; i < 4; i++ {
// 		mutex.Lock()
// 		fmt.Printf("GOROUTINE 1: data=%v, loop=%d\n", data, i+1)
// 		mutex.Unlock()
// 		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
// 	}
// 	done <- true
// }

// func goroutine2(data interface{}, done chan bool) {
// 	for i := 0; i < 4; i++ {
// 		mutex.Lock()
// 		fmt.Printf("GOROUTINE 2: data=%v, loop=%d\n", data, i+1)
// 		mutex.Unlock()
// 		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
// 	}
// 	done <- true
// }

// func main() {
// 	done1 := make(chan bool)
// 	done2 := make(chan bool)

// 	data1 := "hello"
// 	data2 := 42

// 	go goroutine1(data1, done1)
// 	go goroutine2(data2, done2)

// 	<-done1
// 	<-done2

// 	fmt.Println("Both goroutines have finished")
// }

/*

 */
