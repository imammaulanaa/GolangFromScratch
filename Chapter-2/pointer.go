package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	//address1 := Address{"Bogor", "Jawa Barat", "Indonesia"}
	//address2 := address1 // copy value
	//
	//address2.City = "Bekasi"
	//fmt.Println(address1) // tidak berubah
	//fmt.Println(address2) // berubah menjadi Bekasi

	var address1 Address = Address{"Bogor", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // pointer

	address2.City = "Bekasi"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah menjadi Bekasi
}
