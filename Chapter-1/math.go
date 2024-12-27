package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var a = 10
	var b = 10
	var d = 5
	var e = 2
	var c = a + b - d*e
	fmt.Println(c)

	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)

	i += 5 // i = i + 5
	fmt.Println(i)

	var j = 1
	j++ // j = j + 1
	fmt.Println(j)
	j++ // j = j + 1
	fmt.Println(j)

	j--
	fmt.Println(j)
	j--
	fmt.Println(j)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan angka pertama: ")
	input1, _ := reader.ReadString('\n')
	angka1, _ := strconv.ParseFloat(input1[:len(input1)-1], 64)

	fmt.Print("Masukkan angka kedua: ")
	input2, _ := reader.ReadString('\n')
	angka2, _ := strconv.ParseFloat(input2[:len(input2)-1], 64)

	// Melakukan operasi aritmatika
	penjumlahan := angka1 + angka2
	pengurangan := angka1 - angka2
	perkalian := angka1 * angka2
	pembagian := angka1 / angka2

	fmt.Println("Hasil Penjumlahan:", penjumlahan)
	fmt.Println("Hasil Pengurangan:", pengurangan)
	fmt.Println("Hasil Perkalian:", perkalian)
	fmt.Println("Hasil Pembagian:", pembagian)
}
