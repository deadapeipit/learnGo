package main

import (
	"fmt"
	"main/model"
	"os"
	"strconv"
)

func main() {

	//biodata assignment:

	bio1 := model.Biodata{
		Nama:      "Anton",
		Alamat:    "Jakarta",
		Pekerjaan: "Pengangguran",
		Alasan:    "Cepat kaya",
	}

	bio2 := model.Biodata{
		Nama:      "Budi",
		Alamat:    "Bekasi",
		Pekerjaan: "Dukun Santet",
		Alasan:    "Biar bisa santet online",
	}

	bio3 := model.Biodata{
		Nama:      "Cacing",
		Alamat:    "Bogor",
		Pekerjaan: "Bidan",
		Alasan:    "magabut",
	}

	allbio := []model.Biodata{}
	allbio = append(allbio, bio1)
	allbio = append(allbio, bio2)
	allbio = append(allbio, bio3)

	newArgs := os.Args

	if len(newArgs) == 2 {
		l, err := strconv.Atoi(newArgs[1])
		if err != nil {
			fmt.Println("masukin angka!")
			return
		}
		v := allbio[l]
		fmt.Printf("%s\n%s\n%s\n%s\n", v.Nama, v.Alamat, v.Pekerjaan, v.Alasan)

	}

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
