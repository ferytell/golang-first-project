package main

import (
	"fmt"
	"os"
	"strconv"
)

type TemanKelas struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {

	// Data teman kelas
	teman := []TemanKelas{
		{"Andi", "Jakarta", "Mahasiswa", "Mau belajar programming"},
		{"Budi", "Bandung", "Programmer", "Mau tingkatkan skill"},
		{"Cindy", "Surabaya", "Guru", "Butuh pengetahuan tambahan"},
	}
	//teman1 := TemanKelas{"Andi", "Jakarta", "Mahasiswa", "Mau belajar programming"}
	teman2 := TemanKelas{"Budi", "Bandung", "Programmer", "Mau tingkatkan skill"}
	teman3 := TemanKelas{"Cindy", "Surabaya", "Guru", "Butuh pengetahuan tambahan"}

	args := os.Args[1:]

	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("%s is not a valid integer, use a number!\n", arg)
			continue
		}

		if num == 1 {
			printPeople(teman, 1)
			// fmt.Println("Nama\tAlamat\t\tPekerjaan\tAlasan")
			// fmt.Printf("%s\t%s\t\t%s\t%s\n", teman1.Nama, teman1.Alamat, teman1.Pekerjaan, teman1.Alasan)
		} else if num == 2 {
			fmt.Println("Nama\tAlamat\t\tPekerjaan\tAlasan")
			fmt.Printf("%s\t%s\t\t%s\t%s\n", teman2.Nama, teman2.Alamat, teman2.Pekerjaan, teman2.Alasan)
		} else if num == 3 {
			fmt.Println("Nama\tAlamat\t\tPekerjaan\tAlasan")
			fmt.Printf("%s\t%s\t\t%s\t%s\n", teman3.Nama, teman3.Alamat, teman3.Pekerjaan, teman3.Alasan)
		} else {
			fmt.Println("no data found, use other number")
		}
	}
}
func printPeople(teman []TemanKelas, index int) {

	if index < 0 || index >= len(teman) {
		fmt.Println("Invalid index")
		return
	}
	t := teman[index]
	fmt.Printf("%q, %q, %q, %q\n", t.Nama, t.Alamat, t.Pekerjaan, t.Alasan)

	// for _, t := range teman {
	// 	fmt.Println("Nama\tAlamat\t\tPekerjaan\tAlasan")
	// 	fmt.Printf("%s\t%s\t\t%s\t%s\n", t.Nama, t.Alamat, t.Pekerjaan, t.Alasan)
	// }
}

/*
Buatlah sebuah service berupa CLI untuk menampilkan data teman-teman kalian dikelas.
Contohnya, ketika kalian menjalankan perintah

go run biodata.go1

maka data yang akan muncul adalah data teman kalian dengan absen no 1.

Data yang harus ditampilkan yaitu:

	●Nama
	●Alamat
	●Pekerjaan
	●Alasan memilih kelas Golang

Gunakanlah struct dan function untuk menampilkan data tersebut.
*Kalian bisa menggunakan os.Args untuk mendapatkan argument pada terminal.
*/
