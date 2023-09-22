package main

import "fmt"

func luasLingkaran(radius float64) {
	fmt.Println("Radius : ", radius)
	fmt.Println("Area   : ", 3.14*radius*radius, "m^2")

}

func main() {
	fmt.Println("Menghitung luas lingkaran")
	//Isi dengan radius Lingkaran
	luasLingkaran(4)
}
