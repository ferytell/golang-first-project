package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	data1 := []interface{}{"coba1", "coba2", "coba3"}
	data2 := []interface{}{"bisa1", "bisa2", "bisa3"}

	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	go func() {
		for i := 0; i < 4; i++ {
			ch1 <- fmt.Sprint(data1, i+1)
		}
		close(ch1)
		wg.Done()
	}()

	go func() {
		for i := 0; i < 4; i++ {
			ch2 <- fmt.Sprint(data2, i+1)
		}
		close(ch2)
		wg.Done()
	}()

	go func() {
		for {
			select {
			case msg1, ok := <-ch1:
				if !ok {
					return
				}
				fmt.Println(msg1)
			case msg2, ok := <-ch2:
				if !ok {
					return
				}
				fmt.Println(msg2)
			}
		}
	}()

	wg.Wait()
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
