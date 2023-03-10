package main

import (
	"fmt"
)

func main() {

	// input string
	input := "selamat malam"

	createMap(input)
}

func createMap(input string) {

	charCounts := make(map[string]int)

	for _, char := range input {
		fmt.Println(string(char))
		charStr := string(char)
		charCounts[charStr]++
	}

	fmt.Println(charCounts)
}

/*
Buatlah looping dengan variable yang berisi string suatu kalimat
dan pecahlah kalimat tersebut menjadi 1 per 1 kata
Setelah sudah dipecah lakukan perhitungan munculnya kata dari
variable tersebut dengan cara mapping golang

Contoh :
Input : “selamat malam”

output :

s
e
l
a
m
a
t

m
a
l
a
m
map[ :1 a:4 e:1 l:2 m:3 s:1 t:1]
*/
