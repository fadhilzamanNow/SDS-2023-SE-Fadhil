package main

import "fmt"

func luasPersegi(sisi1 float64, sisi2 float64) {
	fmt.Println("Sisi 1 : ", sisi1)
	fmt.Println("Sisi 2 : ", sisi2)
	fmt.Println("Area   : ", sisi1*sisi2, "m^2")
}

func main() {
	fmt.Println("Fungsi yang menghitung area persegi ")
	luasPersegi(32.1231, 10.512)

}
