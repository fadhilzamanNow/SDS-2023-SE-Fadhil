package main

import "fmt"

func main() {

	//membuat variabel array makanan berisi 8 indeks tentang nama nama makanan
	var makanan = [8]string{"Mie", "Sate", "Ayam", "Bubur", "Daging", "Ikan", "Roti", "Permen"}
	//menampilkan 8 indeks array
	fmt.Println(makanan)
	//membuat slice1 yang terdiri dari 5 data
	slice1 := makanan[:5]
	//menampilkan slice1
	fmt.Println(slice1)
	//menambahkan 3 data ke slice1

	//menambahkan data baru ke isi slice1 dimasukkan ke variabel slicebaru
	slicebaru := append(slice1, "KFC", "Hokben", "McD")
	//menampilkan slicebaru
	fmt.Println(slicebaru)
}
