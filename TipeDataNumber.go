package main

import "fmt"

func main() {

	// tipe data integer merupakan tipe data yang berisi angka.
	//memiliki berapa jenis , seperti :
	// int8 min:-128 sampai max:127
	// int16 min:-33.000 sampai max:33.000
	// int32 min:-2miliaran sampai max:2miliaran
	// int64 min:-9triliunan sampai max:9triliunan

	// unsigned integer, bedanya dengan integer nilai minimalnya 0 :
	// uint8 min:0 sampai max:255
	// uint16 min:0 sampai max:67.000
	// uint32 min:0 sampai max:4miliaran
	// uint64 min:0 sampai max:18triliunan

	// tipe data floating point atau float.
	//memiliki beberapa jenis seperti :
	//float32 min:1.18x10^-38 sampai max:3.4x10^38
	//float64 min:2.23x10^-308 sampai max:1.80x10^308

	// Alias atau nama lain tipe data number diatas, seperti :
	// tipe data byte = uint8
	// tipe data rune = int32
	// tipe data int  = minimal int32
	// tipe data uint = minimal uint32

	// penggunaan tipe data integer
	 fmt.Println("ini int/uint :", 1);
	 fmt.Println("ini float :", 1.2);
}
