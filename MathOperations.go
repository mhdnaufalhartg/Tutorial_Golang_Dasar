package main

import "fmt"

func main() {

	// Operasi Matematika di Golang
	// normal operasi : untuk operasi matematika biasa
	// augmented operasi : operasi matematika yang dipermudah
	// unary operasi : operasi matematika berdasarkan kondisi

	// normal operations       // augmented operations    // unary operations
	// + untuk penjumlahan		// += untuk penjumlahan   // ++ untuk a + 1
	// - untuk pengurangan     // -= untuk pengurangan   // -- untuk a - 1
	// * untuk perkalian       // *= untuk perkalian    // + untuk positif
	// / untuk pembagian			// /= untuk pembagian    // = untuk negatif
	// % untuk sisa pembagian  // %= untuk sisa pembagian // ! untuk jika nilai salah

	// contoh penggunaan operasi matematika

	// normal operations
	angka1 := 20
	angka2 := 20

	normalPertambahan := angka1 + angka2;
	normalPengurangan := angka1 - angka2;
	normalPerkalian := angka1 * angka2;
	normalPembagian := angka1 / angka2;
	normalSisaBagi := angka1 % angka2;

	 fmt.Println("Hasil Normal Pertambahan :",normalPertambahan)
	 fmt.Println("Hasil Normal Pengurangan :",normalPengurangan)
	 fmt.Println("Hasil Normal Perkalian :",normalPerkalian)
	 fmt.Println("Hasil Normal Pembagian :",normalPembagian)
	 fmt.Println("Hasil Normal Sisa Pembagian :",normalSisaBagi)

	//  augmented operations
	 angka3 := 30

	 angka3 += 30
	fmt.Println("Hasil Augmented Pertambahan :",angka3)
	angka3 -=30
	fmt.Println("Hasil Augmented Pengurangan :",angka3)
	angka3 *= 30
	fmt.Println("Hasil Augmented Perkalian :",angka3)
	angka3 /= 30
	fmt.Println("Hasil Augmented Pembagian :",angka3)
	angka3 %= 30
	fmt.Println("Hasil Augmented Sisa Pembagian :",angka3)

	// unary operations
	angka4 := 10;

	angka4++
	fmt.Println("Hasil Unary Pertambahan :",angka4)
	angka4--
	fmt.Println("Hasil Unary Pengurangan :",angka4)
	angka5 := +10;
	fmt.Println("Hasil Unary Positif :",angka5)
	angka6 := -10;
	fmt.Println("Hasil Unary Negatif :",angka6)
	if(angka4 != 100){
		fmt.Println("True Unary ! ")
	}else{
		fmt.Println("False Unary ! ")
	}

}
