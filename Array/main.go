package main

import "fmt"

func main() {

	// array dimulai 0
	// array dengan pajang nilai 5
	var a [5]int
	fmt.Println(a) // -> [0,0,0,0,0]

	// mengganti nilai array dengan
	// index 4 (array ke 5) menjadi 1000
	a[4] = 1000
	fmt.Println(a) // -> [0,0,0,0,1000]

	// print array indeks ke 4
	// 	0 1 2 3  4
	// 	| | | |  |
	// [0,0,0,0,1000]
	fmt.Println(a[4])                    // -> 1000
	fmt.Println("panjang array", len(a)) // -> 5

	// contoh deklarasi array dengan panjang 5
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b) // -> [1,2,3,4,5]

	// deklarasi array 2 dimensi
	// array yang memiliki 2 array di dalamnya
	// setiap array panjang 3
	// [[0,0,0] [0,0,0]]
	var twoD [2][3]int
	// mengambahkarn array 2 dimendi
	// pengulangan array dimensi pertama panjang 2
	for i := 0; i < 2; i++ {
		// pengulangan array dimensi kedua
		// setiap array w/ panjang 3
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) // -> [[0 1 2] [1 2 3]]
}
