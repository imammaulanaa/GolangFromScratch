package main

import "fmt"

func main() {
	const (
		firstName string = "Imam"
		lastName         = "Maulana"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	//error
	//firstName = "akbar"
	//lastName = "Imam"
}
