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
