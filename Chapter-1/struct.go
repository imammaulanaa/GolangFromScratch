package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hi", name, "my name is", customer.Name)
}

func main() {
	var imam Customer
	fmt.Println(imam)

	imam.Name = "Imam Maulana"
	imam.Address = "Bekasi"
	imam.Age = 30
	fmt.Println(imam)
	fmt.Println(imam.Name)
	fmt.Println(imam.Address)
	fmt.Println(imam.Age)

	// struct literal
	akbar := Customer{
		Name:    "akbar",
		Address: "Jakarta",
		Age:     30,
	}
	fmt.Println(akbar)

	agus := Customer{"agus", "Depok", 30}
	fmt.Println(agus)

	akbar.sayHi("imam")
	imam.sayHi("akbar")
	agus.sayHi("Agus")
}
