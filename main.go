package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printData(data interface{}, num int, c chan int) {
	for i := 0; i < 4; i++ {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Printf("Data %d: %v, Loop %d\n", num, data, i+1)
	}
	c <- num
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	data1 := []interface{}{"foo", 123, true}
	data2 := []interface{}{"bar", 456, false}

	go func() {
		for i := 0; i < 1; i++ {
			randIndex := rand.Intn(len(data1))
			go printData(data1[randIndex], 1, c1)
		}
	}()

	go func() {
		for i := 0; i < 1; i++ {
			randIndex := rand.Intn(len(data2))
			go printData(data2[randIndex], 2, c2)
		}
	}()

	for i := 0; i < 8; i++ {
		select {
		case <-c1:
			fmt.Println("Data 1 done")
		case <-c2:
			fmt.Println("Data 2 done")
		}
	}
}

/*

 */
