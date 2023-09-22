package main

import "fmt"

func luasSegitiga(alas float64, tinggi float64) {
	fmt.Println("Alas   : ", alas)
	fmt.Println("Tinggi : ", tinggi)
	fmt.Println("Area   : ", (alas*tinggi)/2, "m^2")
}

func main() {
	fmt.Println("Fungsi yang menghitung area segitiga ")
	luasSegitiga(6, 2)

}
