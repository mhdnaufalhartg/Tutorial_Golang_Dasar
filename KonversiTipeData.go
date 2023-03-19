package main

import "fmt"

func main() {

	// konversi tipe data adalah mengkonversi value dari si tipe data.
	// Namun perhatikan tipe data yang mau di konversikan dapat menampung nilai dari si data. Jika tidak maka hasilnya akan berubah. Contoh : int64 di konversi ke int64.
	// contoh penggunaan Konversi tipe data

	var abc int64 = 10000;
	abcd := int8(abc)

	fmt.Println(abc)
	 fmt.Println(abcd)
}
