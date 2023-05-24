package main

import "fmt"

func main() {
	var diagonal1, diagonal2, luas int

	fmt.Println("Masukkan panjang diagonal1: ")
	fmt.Scan(&diagonal1)

	fmt.Println("Masukkan panjang diagonal2: ")
	fmt.Scan(&diagonal2)

	luas = diagonal1 * diagonal2
	fmt.Println("Hasil = ", luas)
}
