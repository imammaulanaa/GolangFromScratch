package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	imam := Man{"imam"}
	imam.Married()

	fmt.Println(imam.Name)
}
