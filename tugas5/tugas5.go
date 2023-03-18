package tugas5

import (
	"fmt"
)

func MyGoroutine(data []interface{}, ch chan<- interface{}) {
	//defer close(done)
	for i := 0; i < 4; i++ {
		ch <- fmt.Sprint(data, i+1)
	}
	close(ch)
}

// func Goroutine1(data1 []interface{}) {
// 	wg := sync.WaitGroup{}
// 	wg.Add(2)

// 	for i := 0; i < 4; i++ {
// 		for _, d := range data1 {
// 			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
// 			fmt.Println(d, i+1)
// 		}
// 	}
// 	wg.Done()
// }

// func Goroutine2(data2 []interface{}) {
// 	wg := sync.WaitGroup{}
// 	wg.Add(2)
// 	for i := 0; i < 4; i++ {
// 		for _, d := range data2 {
// 			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
// 			fmt.Println(d, i+1)
// 		}
// 	}
// 	wg.Done()
// }

//result := []string{}

// func Goroutine1(data1 []string, genChan chan string) {
// 	//fmt.Println("test1")
// 	for _, s := range data1 {
// 		//fmt.Println("test3")
// 		genChan <- s
// 	}
// 	close(genChan)
// 	//fmt.Println("test2")
// }

// func Goroutine2(data interface{}, done chan bool) {
// 	for i := 0; i < 4; i++ {
// 		fmt.Printf("[%v, loop=%d] 2\n", data, i+1)
// 		//		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
// 	}
// 	done <- true

// }

//=================================================
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	wg := sync.WaitGroup{}
// 	wg.Add(2)

// 	data1 := []interface{}{"coba1", "coba2", "coba3"}
// 	data2 := []interface{}{"bisa1", "bisa2", "bisa3"}

// 	ch1 := make(chan interface{}, 1)
// 	ch2 := make(chan interface{}, 1)

// 	var mutex sync.Mutex

// 	go func() {
// 		for i := 0; i < 4; i++ {
// 			ch1 <- fmt.Sprint(data1, i+1)
// 		}
// 		close(ch1)
// 		wg.Done()
// 	}()

// 	go func() {
// 		for i := 0; i < 4; i++ {
// 			ch2 <- fmt.Sprint(data2, i+1)
// 		}
// 		close(ch2)
// 		wg.Done()
// 	}()

// 	go func() {
// 		for {
// 			select {
// 			case msg1, ok := <-ch1:
// 				if !ok {
// 					return
// 				}
// 				mutex.Lock()
// 				fmt.Println(msg1)
// 				mutex.Unlock()
// 			case msg2, ok := <-ch2:
// 				if !ok {
// 					return
// 				}
// 				mutex.Lock()
// 				fmt.Println(msg2)
// 				mutex.Unlock()
// 			}
// 		}
// 	}()

// 	wg.Wait()
// }

//============================================================================

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	// Create a WaitGroup with a counter of 1
// 	wg := sync.WaitGroup{}
// 	wg.Add(1)

// 	// Create a buffered channel of interfaces with a capacity of 2
// 	ch := make(chan interface{}, 2)

// 	// Create a mutex
// 	var mutex sync.Mutex

// 	// Create a single goroutine that sends messages to ch
// 	go func() {
// 		defer wg.Done() // Signal that the goroutine has finished

// 		data1 := []interface{}{"coba1", "coba2", "coba3"}
// 		data2 := []interface{}{"bisa1", "bisa2", "bisa3"}

// 		for i := 0; i < 4; i++ {
// 			if i%2 == 0 {
// 				ch <- fmt.Sprint(data1, i+1)
// 			} else {
// 				ch <- fmt.Sprint(data2, i+1)
// 			}
// 		}

// 		for j := 0; j < 4; j++ {
// 			if j%2 == 0 {
// 				ch <- fmt.Sprint(data1, j+1)
// 			} else {
// 				ch <- fmt.Sprint(data2, j+1)
// 			}
// 		}

// 		close(ch)
// 	}()

// 	// Receive messages from ch and print them to stdout
// 	for msg := range ch {
// 		mutex.Lock()
// 		fmt.Println(msg)
// 		mutex.Unlock()
// 	}

// 	// Wait for the goroutine to finish
// 	wg.Wait()
// }
