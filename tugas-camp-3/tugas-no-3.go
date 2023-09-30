package main

import "fmt"

type Pengusaha struct {
	Nama       string
	Kekayaan   string
	Perusahaan string
	Umur       int
}

func main() {
	//manipulasi struct pada property Kekayaan
	fadhil := Pengusaha{
		Nama:       "Muhammad Ilham Isfadhillah",
		Kekayaan:   "$99999999",
		Perusahaan: "Tesla",
		Umur:       25,
	}

	fmt.Println("Kekayaan dari", fadhil.Nama, "pada tahun 2027 sebanyak", fadhil.Kekayaan)
	fadhil.Kekayaan = "$10000000000000000000000000000000"
	fmt.Println("Namun pada 2028, kekayaannya bertambah menjadi", fadhil.Kekayaan)

	//manipulasi pointer dengan memanipulasi variabel makananFavorit yang akan ditampilkan dengan medereference makananFavorit1
	var makananFavorit string = "Mie Goreng"

	var makananFavorit1 *string = &makananFavorit
	fmt.Println("Makanan Favoritku kemarin adalah", *makananFavorit1, "[ alamat : ", makananFavorit1, "]")
	makananFavorit = "Nasi Padang"
	fmt.Println("Makanan Favoritku hari ini adalah", *makananFavorit1, "[ alamat : ", makananFavorit1, "]")

}
