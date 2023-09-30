package main

import "fmt"

func main() {

	infoOrang1 := map[string]string{
		"nama":       "Bejo Unsoed",
		"pendidikan": "S1",
		"asal":       "Bandung",
		"hobi":       "Coding",
	}

	infoOrang2 := map[string]string{
		"nama":       "Jarwo Unsoed",
		"pendidikan": "S3",
		"asal":       "Bandung",
		"hobi":       "Melakukan Penelitian",
	}

	infoOrang3 := map[string]string{
		"nama":       "Sopo Unsoed",
		"pendidikan": "S3",
		"asal":       "Bandung",
		"hobi":       "Melakukan Eksperimen",
	}

	fmt.Println(infoOrang1)
	fmt.Println("Nama :", infoOrang1["nama"])
	fmt.Println("Pendidikan :", infoOrang1["pendidikan"])
	fmt.Println("Asal :", infoOrang1["asal"])
	fmt.Println("Hobi :", infoOrang1["hobi"])

	fmt.Println()
	fmt.Println(infoOrang2)
	fmt.Println("Nama :", infoOrang2["nama"])
	fmt.Println("Pendidikan :", infoOrang2["pendidikan"])
	fmt.Println("Asal :", infoOrang2["asal"])
	fmt.Println("Hobi :", infoOrang2["hobi"])

	fmt.Println()
	fmt.Println(infoOrang3)
	fmt.Println("Nama :", infoOrang3["nama"])
	fmt.Println("Pendidikan :", infoOrang3["pendidikan"])
	fmt.Println("Asal :", infoOrang3["asal"])
	fmt.Println("Hobi :", infoOrang3["hobi"])
}
