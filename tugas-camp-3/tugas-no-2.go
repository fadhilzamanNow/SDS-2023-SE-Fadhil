package main

import "fmt"

type Pengusaha struct {
	Nama       string
	Kekayaan   string
	Perusahaan string
	Umur       int
}

func (s Pengusaha) ucapSalam() {
	fmt.Println("Selamat Siang Bapak,", s.Nama)
}

func (s Pengusaha) tanya() {
	fmt.Println("Apakah bapak membuka lowongan perkerjaan di", s.Perusahaan)
}

func main() {
	fadhil := Pengusaha{
		Nama:       "Muhammad Ilham Isfadhillah",
		Kekayaan:   "$99999999",
		Perusahaan: "Tesla",
		Umur:       25,
	}

	fmt.Println("TRENDING TOPIC 2027 ")
	fmt.Println("Dikutip dari detiknews.com, pengusaha bernama", fadhil.Nama, "yang merupakan pengusaha baru yang akhir akhir ini hangat namanya dari perusahaan", fadhil.Perusahaan)
	fmt.Println("dia adalah pengusaha termuda yang berhasil mengumpulkan kekayaan sejumlah", fadhil.Kekayaan, "di umurnya yang masih", fadhil.Umur)
	fmt.Println()
	fadhil.ucapSalam()
	fadhil.tanya()
}
