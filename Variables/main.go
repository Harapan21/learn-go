package main

import "fmt"

func main() {
	// menggunakan var teknik pertama
	var a = "this is var"
	fmt.Println(a)

	// menggunakan var teknik kedua
	// var sekaligus deklarasi dua, dan type datanya
	var v, c int = 1, 2
	fmt.Println(v, c)

	// deklarasi var teknik ketiga
	var d = true
	fmt.Println(d)

	// deklarasi var teknik keempat
	// var dengan tipe tipe data integer, dengan nilai kosong
	// otomatis akan ke nilai 0
	var e int
	fmt.Println(e)

	// deklarasi variable dengan cara cepat tanpa menggunakan
	// tipe data
	// contoh dibawah sama saja dengan
	// var f string = "cara cepat"
	f := "cara cepat"
	fmt.Println(f)
}
