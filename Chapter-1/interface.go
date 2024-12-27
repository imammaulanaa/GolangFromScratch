package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHi(value HasName) {
	fmt.Println("Hi", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Imam"}
	SayHi(person)

	animal := Animal{Name: "Kucing"}
	SayHi(animal)
}
