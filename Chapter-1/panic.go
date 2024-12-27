package main

import "fmt"

func endApp() {
	fmt.Println("End app")
	message := recover()
	fmt.Println("system panic", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error")
	}
}

func main() {
	runApp(true)
	fmt.Println("Jalan lagi setelah panic")
}
