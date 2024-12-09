package main

import "fmt"

func main() {
	//var person map[string]string = map[string]string{}
	//person["name"] = "Imam"
	//person["address"] = "Bekasi"

	person := map[string]string{
		"name":    "Imam",
		"address": "Bekasi",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	git := make(map[string]string)
	git["title"] = "GolangFromScratch"
	git["author"] = "Imam"
	git["ups"] = "Salah"

	fmt.Println(git)

	delete(git, "ups")

	fmt.Println(git)
}
