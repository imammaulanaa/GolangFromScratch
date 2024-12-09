package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Membuat objek reader untuk membaca input dari pengguna.
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Halo! Siapa namamu? ")
	// Membaca input nama dari pengguna hingga menemukan baris baru.
	nama, _ := reader.ReadString('\n')
	// Menghilangkan karakter baris baru dari input nama.
	nama = nama[:len(nama)-1]

	fmt.Println("Senang bertemu denganmu,", nama+"!")
}
