package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Println(arr)
			fmt.Println("Elemen yang ditukar yaitu ", arr[j], " dengan ", arr[j+1])
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			} else {
				fmt.Println("tidak ada penukaran")
			}
		}
	}
}

func main() {
	data := []int{9, 4, 5, 6, 7, 2, 32, 1}
	//bentuk awal array
	fmt.Println("BENTUK SEBELUM DITUKAR : ", data)
	//array diproses diurutkan dari yang tinggi ke rendah
	bubbleSort(data)
	//bentuk akhir array
	fmt.Println("BENTUK SETELAH DITUKAR : ", data)
}
