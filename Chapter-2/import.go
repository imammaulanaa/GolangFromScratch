package main

import (
	"fmt"
	"chapter2/helper"
)

func main() {
	result := helper.Sayhi("Imam")
	fmt.Println(result)

	// fmt.Println(helper.Application)
	// fmt.Println(helper.version) // tidak bisa diakses
	// fmt.Println(helper.sayGoodBye("Imam")) // tidak bisa diakses
}
