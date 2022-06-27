package main

import (
	"fmt"
	"main/model"
	"os"
	"strconv"
)

func printBio(i []string, data []model.Biodata) {
	if len(i) == 2 {
		l, err := strconv.Atoi(i[1])
		if err != nil {
			fmt.Println("masukin angka!")
			return
		}
		v := data[l-1]
		fmt.Printf("%s\n%s\n%s\n%s\n", v.Nama, v.Alamat, v.Pekerjaan, v.Alasan)
	} else {
		fmt.Println("belum masukin angka bos!")
	}
}

func main() {

	//biodata assignment:

	allbio := []model.Biodata{
		{
			Nama:      "Anton",
			Alamat:    "Jakarta",
			Pekerjaan: "Pengangguran",
			Alasan:    "Cepat kaya",
		}, {
			Nama:      "Budi",
			Alamat:    "Bekasi",
			Pekerjaan: "Dukun Santet",
			Alasan:    "Biar bisa santet online",
		}, {
			Nama:      "Cacing",
			Alamat:    "Bogor",
			Pekerjaan: "Bidan",
			Alasan:    "magabut",
		},
	}

	newArgs := os.Args

	printBio(newArgs, allbio)

	// //slice:
	// var slice = []string{"Budi", "Cacing", "Dede", "Gemez"}
	// for _, item := range slice {
	// 	fmt.Println(item)
	// }

	//looping
	// for j := 0; j < 11; j++ {
	// 	if j%2 == 1 {
	// 		fmt.Println("ganjil")

	// 	} else {
	// 		fmt.Println("genap")

	// 	}
	// }

	//Hello world:
	//fmt.Println("Hello world")
}
