package main

import (
	"fmt"
	"os"
)

type TemanKelas struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {

	fmt.Println("Nama program:", os.Args[0])
	fmt.Println("Argumen baris perintah:", os.Args[1])
	// Data teman kelas
	teman1 := TemanKelas{"Andi", "Jakarta", "Mahasiswa", "Mau belajar programming"}
	teman2 := TemanKelas{"Budi", "Bandung", "Programmer", "Mau tingkatkan skill"}
	teman3 := TemanKelas{"Cindy", "Surabaya", "Guru", "Butuh pengetahuan tambahan"}

	// Menampilkan data teman kelas
	fmt.Println("Data Teman Kelas")
	fmt.Println("Nama\tAlamat\t\tPekerjaan\tAlasan", os.Args[2])
	fmt.Printf("%s\t%s\t%s\t\t%s\n", teman1.Nama, teman1.Alamat, teman1.Pekerjaan, teman1.Alasan)
	fmt.Printf("%s\t%s\t%s\t%s\n", teman2.Nama, teman2.Alamat, teman2.Pekerjaan, teman2.Alasan)
	fmt.Printf("%s\t%s\t%s\t\t%s\n", teman3.Nama, teman3.Alamat, teman3.Pekerjaan, teman3.Alasan)

}

/*
Buatlah sebuah service berupa CLI untuk menampilkan data teman-teman kalian dikelas.
Contohnya, ketika kalian menjalankan perintah go run biodata.go
maka data yang akan muncul adalah data teman kalian dengan absen no 1.

Data yang harus ditampilkan yaitu:

	●Nama
	●Alamat
	●Pekerjaan
	●Alasan memilih kelas Golang

Gunakanlah struct dan function untuk menampilkan data tersebut.
*Kalian bisa menggunakan os.Args untuk mendapatkan argument pada terminal.
*/
