package main

import "fmt"

// Fungsi untuk mengkonversi suhu dari Celsius ke Fahrenheit
func celsiusKeFahrenheit(celsius float64) {
	fmt.Println(celsius, "Celsius =", (celsius*9/5)+32, "Fahrenheit")
}

// Fungsi untuk mengkonversi suhu dari Fahrenheit ke Celsius
func fahrenheitKeCelsius(fahrenheit float64) {
	fmt.Println(fahrenheit, "Fahrenheit =", (fahrenheit-32)*5/9, "Celsius")

}

// // Fungsi untuk mengkonversi suhu dari Celsius ke Kelvin
func celsiusKeKelvin(celsius float64) {
	fmt.Println(celsius, "%.2f Celsius =", celsius+273.15, "Kelvin")

}

// // Fungsi untuk mengkonversi suhu dari Kelvin ke Celsius
func kelvinKeCelsius(kelvin float64) {
	fmt.Println(kelvin, "Kelvin =", kelvin-273.15, "Celsius")

}

// // Fungsi untuk mengkonversi suhu dari Celsius ke Réaumur
func celsiusKeReaumur(celsius float64) {
	fmt.Println(celsius, "Celsius =", celsius*4/5, "Réaumur")
}

// // Fungsi untuk mengkonversi suhu dari Réaumur ke Celsius
func reaumurKeCelsius(reaumur float64) {
	fmt.Println(reaumur, "Réaumur =", reaumur*5/4, "Celsius")
}

func main() {
	fmt.Println("Konversi Suhu")

	// Konversi dari Celsius ke Fahrenheit
	celsiusKeFahrenheit(30)

	// Konversi dari Fahrenheit ke Celsius

	fahrenheitKeCelsius(45)

	// // Konversi dari Celsius ke Kelvin

	celsiusKeKelvin(50)

	// // Konversi dari Kelvin ke Celsius

	kelvinKeCelsius(400)

	// // Konversi dari Celsius ke Réaumur

	celsiusKeReaumur(80)

	// // Konversi dari Réaumur ke Celsius

	reaumurKeCelsius(40)

}
