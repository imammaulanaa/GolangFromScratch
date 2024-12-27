package main

import "fmt"

func main() {
	name := "Imam Maulana"
	fmt.Println(name)

	name = "Imam Akbar"
	fmt.Println(name)

	var (
		firstName  = "Imam"
		middleName = "Maulana"
		lastName   = "Akbar"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
