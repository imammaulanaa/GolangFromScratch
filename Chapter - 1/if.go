package main

import "fmt"

func main() {
	name := "Imam"

	if name == "Imam" {
		fmt.Println("Hi Imam")
	} else if name == "akbar" {
		fmt.Println("Hi Akbar")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
