package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Bogor", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // pointer
	address2.City = "Bekasi"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah menjadi Bekasi

	//address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}
