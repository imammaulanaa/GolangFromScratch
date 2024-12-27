package main

import "fmt"

func main() {
	type NoKTP string

	var ktpimam NoKTP = "123456789"

	var contoh string = "987654321"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpimam)
	fmt.Println(contohKtp)
}
