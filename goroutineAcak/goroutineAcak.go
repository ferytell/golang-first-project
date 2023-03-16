package goroutineacak

import (
	"fmt"
)

func Goroutine1(data interface{}, done chan bool) {
	for i := 0; i < 4; i++ {
		fmt.Printf("[%v, loop=%d] 1\n", data, i+1)
		//		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
	done <- true

}

func Goroutine2(data interface{}, done chan bool) {
	for i := 0; i < 4; i++ {
		fmt.Printf("[%v, loop=%d] 2\n", data, i+1)
		//		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
	done <- true

}
