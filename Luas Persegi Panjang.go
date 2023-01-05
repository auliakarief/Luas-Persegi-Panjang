package main

import "fmt"

func main() {
	fmt.Println("HALO")

	var panjang, lebar, luas int

	// Input Panjang
	fmt.Println("Masukkan Panjang!")
	fmt.Scan(&panjang)

	// Input Lebar
	fmt.Println("Masukkan Lebar!")
	fmt.Scan(&lebar)

	luas = panjang * lebar

	fmt.Println("-----------------------")
	fmt.Println("Panjang  : ", panjang)
	fmt.Println("Lebar  : ", lebar)
	fmt.Print("Luas adalah  : ", luas)
}
