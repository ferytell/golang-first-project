package main

import (
	"fmt"
	"strconv"
)

func main() {

	var i int = 8
	x := "%"
	var j bool = true
	k := 123.456
	s := strconv.FormatFloat(k, 'f', -1, 64)

	numbase10 := 21
	numbase8 := 25
	numbase16 := "f"
	numbase16_2 := "F 13"
	ya := 'Я'
	flo := 123.456000
	floSe := 1.23456

	/*	const (
		c1 = iota
		c2
		c3
	)		*/

	fmt.Println(x)
	fmt.Printf("tipe data i: %T \n", i)
	fmt.Println(x, "\n", j)
	fmt.Println("\u042F", "(ya)")
	fmt.Printf("nilai base 10 : 21 adalah %d \n", numbase10)
	fmt.Printf("nilai base 8 : 25 adalah %o \n", numbase8)
	fmt.Printf("nilai base 16 : f adalah %x \n", numbase16)
	fmt.Printf("nilai base 16 : F 13 adalah %x \n", numbase16_2)
	fmt.Printf("unicode dari Я %U \n ", ya)
	fmt.Println("unicode dari var k")
	for _, c := range s {
		fmt.Printf("%U ", c)
	}
	fmt.Printf("\n %f \n", flo)
	fmt.Printf(" %e \n", floSe)

}

/*
menampilkan nilai i : 21 fmt.Printf("%T \n", i)
// contoh : fmt.Printf("%v \n", i)

menampilkan tipe data dari variabel i
menampilkan tanda %
menampilkan nilai boolean j : true
menampilkan nilai boolean j : true
menampilkan unicode russia : Я (ya)
menampilkan nilai base 10 : 21
menampilkan nilai base 8 :25
menampilkan nilai base 16 : f
menampilkan nilai base 16 : F 13
menampilkan unicode karakter Я : U+042F var k float64 = 123.456;
menampilkan float : 123.456000
menampilkan float scientific : 1.234560E+02
*/
