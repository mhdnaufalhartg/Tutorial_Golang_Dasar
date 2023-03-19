package main

import "fmt"

func main() {

	// Operasi Boolean untuk cek apakah data bernilai false atau true

	// Contoh Penggunaan Operasi Boolean :

	angka1 := 10
	angka2 := 10

	// Cek apakah angka1>5 dan apakah angka2>9?
	cekAngka1 := angka1 > 5
	cekAngka2 := angka2 > 9

	HasilAngka1 := cekAngka1
	Hasilangka2 := cekAngka2

	// DAN
	result := HasilAngka1 && Hasilangka2
	fmt.Println(result)

	// ATAU
	result2 := HasilAngka1 || Hasilangka2
	fmt.Println(result2)

	// TIDAK SAMA
	result3 := HasilAngka1 && Hasilangka2
	fmt.Println(!result3)


}
