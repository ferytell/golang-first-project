package main

import "fmt"

func main() {

	var i int = 8
	x := "%"
	var j bool = true

	const (
		c1 = iota
		c2
		c3
	)
	fmt.Println(x, c2, c3)
	fmt.Printf("tipe data i: %T \n", i)
	fmt.Println(x, "\n", j)

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
