package main

import "fmt"

func main() {
	// Deklarasi variabel beserta tipe data
	var nama string = "rahman";
	var umur int = 23;

	fmt.Printf("nama : %s, umur : %d \n", nama, umur);

	// hanya deklarasi
	var alamat string;
	alamat = "sukabumi";

	fmt.Printf("alamat : %s \n", alamat);

	// type inference langsung diisi data nya
	var status = "lajang";
	fmt.Println("status :"+ status);

	// deklarasi tanpa tipe data var dan diisi langsung datanya
	lulusan := "Universitas Muhammadiya Sukabumi";
	fmt.Println(lulusan);

	// Multi variabel
	var satu, dua, tiga string = "satu", "dua", "tiga";
	fmt.Println(satu, dua, tiga);

	// dengan :=
	one, two, three := "one", 2, "three";
	fmt.Println(one, two, three);

	// variable underscore _ menampung nilai yang tidak pakai
	// jika dalam go variabel yang ada tapi tidak dipakai akan error
	// maka dari itu bisa disimpan di _ tapi dipanggil akan error
	// _ hanya sebagai keranjang sampah
	name, _ := "lagi", "Belajar Golang";
	fmt.Println(name);


	// variabel pointer +>mereference ke alamat memori
	name1 := new(string);
	fmt.Println(name1);

	
}
