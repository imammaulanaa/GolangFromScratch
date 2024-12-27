package main

import (
	"fmt"
	"chapter2/database"
	_ "chapter2/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}
