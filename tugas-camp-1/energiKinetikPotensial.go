package main

import "fmt"

func Kinetik(massa float64, kecepatan float64) {
	fmt.Println("Dengan kecepatan", kecepatan, " m/s dan massa dari benda sebesar", massa, "kg maka energi kinetik yang dihasilkan adalah :", 0.5*massa*kecepatan*kecepatan, "Joule")
}

func Potensial(massa float64, tinggi float64) {
	fmt.Println("Dengan mas,sa", massa, "kg dan tinggi antara tanah dengan benda", tinggi, "m dan gravitasi di bumi sebesar 9,8 m/s^2 maka energi potensial yang dihasilkan adalah :", massa*9.8*tinggi, "Joule")
}

func main() {
	fmt.Println("Contoh Energi Kinetik : ")
	Kinetik(2, 4)
	Potensial(5, 10)
}
