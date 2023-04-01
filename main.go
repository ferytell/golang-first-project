package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	for i := 0; i < 4; i++ {
		func(j int) {
			go async1(j)
			go async2(j)
		}(i)
	}
	// func() {
	// 	for i := 0; i < 4; i++ {
	// 		go async1(i + 1)
	// 		go async2(i + 1)
	// 	}
	// }()
	time.Sleep(2 * time.Second)
	fmt.Println("space--------")

	var mx sync.Mutex
	for i := 0; i < 4; i++ {
		go func(j int) {
			mx.Lock()
			//mx.Unlock()
			async1(j)
			//mx.Lock()
			async2(j)
			mx.Unlock()
			time.Sleep(time.Second)
		}(i)
	}
	time.Sleep(2 * time.Second)

}

func async1(i int) {
	data := []interface{}{"coba1", "coba2", "coba3"}
	fmt.Println(data, i)
}

func async2(i int) {
	data := []interface{}{"bisa1", "bisa2", "bisa3"}
	fmt.Println(data, i)
}

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wg = sync.WaitGroup{}
// var mutex sync.Mutex

// func main() {
// 	ch1 := make(chan interface{}, 20)
// 	// ch2 := make(chan interface{}, 5)
// 	// ch3 := make(chan interface{}, 5)

// 	data1 := []interface{}{"coba1", "coba2", "coba3"}
// 	data2 := []interface{}{"bisa1", "bisa2", "bisa3"}

// 	// Create a goroutine that sends messages to channel

// 	for i := 0; i < 4; i++ {
// 		go func(j int) {
// 			wg.Add(2)
// 			//mutex.Lock()
// 			goroutineAcak(j, ch1, data1, &wg)
// 			goroutineAcak(j, ch1, data2, &wg)
// 			//mutex.Unlock()

// 		}(i)
// 	}
// 	//time.Sleep(1 * time.Second)
// 	fmt.Println("space--------")

// 	for i := 0; i < 4; i++ {
// 		go func(j int) {
// 			wg.Add(2)
// 			//mutex.Lock()
// 			goroutineRapi(j, ch1, data1, &wg)
// 			goroutineRapi(j, ch1, data2, &wg)
// 			//mutex.Unlock()

// 		}(i)
// 	}

// 	wg.Wait()

// 	close(ch1)
// 	// Receive messages from ch and print them to stdout

// 	// for i := 0; i < 4; i++ {
// 	// 	go pleaseWorks(i, ch2, data1, data2)
// 	// 	go pleaseWorks(i, ch3, data1, data2)

// 	// }

// 	for item := range ch1 {
// 		//mutex.Lock()
// 		fmt.Println(item)
// 		//mutex.Unlock()
// 	}

// }

// func goroutineAcak(i int, ch chan<- interface{}, data1 []interface{}, wg *sync.WaitGroup) {
// 	ch <- fmt.Sprint(data1, i+1)
// 	//wg.Done()
// }

// func goroutineRapi(i int, ch chan<- interface{}, data1 []interface{}, wg *sync.WaitGroup) {
// 	mutex.Lock()
// 	ch <- fmt.Sprint(data1, i+1)
// 	mutex.Unlock()
// 	wg.Done()
// }

// // func pleaseWorks(i int, ch chan<- interface{}, data1 []interface{}, data2 []interface{}) {
// // 	defer wg.Done() // Signal that the goroutine has finished
// // 	mutex.Lock()
// // 	if i%2 == 0 {
// // 		ch <- fmt.Sprint(data1, i/2+1)
// // 	} else {
// // 		ch <- fmt.Sprint(data2, i/2+1)
// // 	}
// // 	mutex.Unlock()
// // }
