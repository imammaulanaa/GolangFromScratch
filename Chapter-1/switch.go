package main

import "fmt"

func main() {
	name := "Imam"

	switch name {
	case "Imam":
		fmt.Println("Hi, Imam")
	case "Akbar":
		fmt.Println("Hi, Akbar")
	default:
		fmt.Println("Hi, Boleh kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	name = "Imam Maulana"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
