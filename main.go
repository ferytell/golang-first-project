package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var mutex sync.Mutex

func main() {
	ch1 := make(chan interface{}, 20)

	data1 := []interface{}{"coba1", "coba2", "coba3"}
	data2 := []interface{}{"bisa1", "bisa2", "bisa3"}

	// Create a goroutine that sends messages to channel
	for i := 0; i < 4; i++ {
		wg.Add(3)
		go pleaseWork(i, ch1, data1)
		go pleaseWork(i, ch1, data2)
		go pleaseWorks(i, ch1, data1, data2)
	}

	wg.Wait()
	close(ch1)
	// Receive messages from ch and print them to stdout

	for item := range ch1 {
		//mutex.Lock()
		fmt.Println(item)
		//mutex.Unlock()
	}

	// for msg := range ch1 {
	// 	//mutex.Lock()
	// 	fmt.Println(msg)
	// 	//mutex.Unlock()
	// }
	// for msg := range ch2 {
	// 	//mutex.Lock()
	// 	fmt.Println(msg)
	// 	//mutex.Unlock()
	// }

}

func pleaseWork(i int, ch chan<- interface{}, data []interface{}) {
	defer wg.Done() // Signal that the goroutine has finished
	mutex.Lock()
	ch <- fmt.Sprint(data, i+1)
	mutex.Unlock()
}

func pleaseWorks(i int, ch chan<- interface{}, data1 []interface{}, data2 []interface{}) {
	defer wg.Done() // Signal that the goroutine has finished
	mutex.Lock()
	if i%2 == 0 {
		ch <- fmt.Sprint(data1, i/2+1)
	} else {
		ch <- fmt.Sprint(data2, i/2+1)
	}
	mutex.Unlock()
}

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// func main() {
// 	wg := sync.WaitGroup{}
// 	wg.Add(2)

// 	data1 := []interface{}{"coba1", "coba2", "coba3"}
// 	data2 := []interface{}{"bisa1", "bisa2", "bisa3"}

// 	mu1 := sync.Mutex{}
// 	mu2 := sync.Mutex{}

// 	go func() {
// 		for i := 0; i < 4; i++ {
// 			mu1.Lock()
// 			for _, d := range data1 {
// 				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
// 				fmt.Println(d, i+1)
// 			}
// 			mu1.Unlock()
// 		}
// 		wg.Done()
// 	}()

// 	go func() {
// 		for i := 0; i < 4; i++ {
// 			mu2.Lock()
// 			for _, d := range data2 {
// 				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
// 				fmt.Println(d, i+1)
// 			}
// 			mu2.Unlock()
// 		}
// 		wg.Done()
// 	}()

// 	wg.Wait()
// }
