package main

import "fmt"

func main() {
	kondisi1 := 10 > 5  // Benar
	kondisi2 := 20 < 30 // Benar

	keduakondisibenar := kondisi1 && kondisi2

	fmt.Println("Apakah kedua kondisi benar? ", keduakondisibenar)
}
