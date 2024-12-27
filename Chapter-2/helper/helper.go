package helper

import "fmt"

var version = "1.0.0"
var Application = "golang"

func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func Sayhi(name string) string {
	return "Hello " + name
}

func Contoh() {
	sayGoodBye("Imam")
	fmt.Println(version)
}