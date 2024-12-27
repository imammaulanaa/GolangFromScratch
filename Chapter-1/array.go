package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Imam"
	names[1] = "Maulana"
	names[2] = "Akbar"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{
		90,
		80,
		95,
		100,
		70,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[4])
	fmt.Println(values[2])

	fmt.Println(len(values))
}
