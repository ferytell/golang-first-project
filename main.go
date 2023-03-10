package main

import (
	"fmt"
)

func main() {
	i := 0
	j := 0
	char := []string{"С", "А", "Ш", "А", "Р", "В", "О"}

	//fmt.Println(num)
	for i < 6 {
		fmt.Println("Nilai i =", i)
		i++
		if i == 5 {
			for j < 11 {
				fmt.Println("Nilai j =", j)
				j++
				if j == 5 {
					printChar(char)
				}
			}
			break
		}
	}
}

func printChar(char []string) {
	x := 0
	for _, s := range char {
		r := []rune(s)[0]
		fmt.Printf("character %U", r)
		fmt.Println(" '", s, "' start at byte position", x)
		x += 2

	}
}

/*
Nilai i = 0
Nilai i = 1
Nilai i = 2
Nilai i = 3
Nilai i = 4
Nilai j = 0
Nilai j = 1
Nilai j = 2
Nilai j = 3
Nilai j = 4
character U+0421 'C' start at byte position 0
character U+0410 'A' start at byte position 2
character U+0428 'Ш' start at byte position 4
character U+0410 'A' start at byte position 6
character U+0420 'P' start at byte position 8
character U+0412 'B' start at byte position 10
character U+041E 'O' start at byte position 12
Nilai j = 6
Nilai j = 7
Nilai j = 8
Nilai j = 9
Nilai j = 10

*/
