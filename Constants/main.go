package main

import (
	"fmt"
	"math"
)

// constant (konstanta)
// nilai yang hanya di tentukan sekali
// lebih lanjut
// https://id.wikipedia.org/wiki/Konstanta_(pemrograman)

const s string = "constant"

func main() {
	// s = "tidak bisa diubah"
	// var f string  = "bisa di ubah"
	// f = "hallo"
	fmt.Println(s)

	const n = 500000000
	// konstanta proses aritmatika
	// bisa "dikelola sekalo"
	// tetapi hasil dari nilai arimatika tidak bisa dirubah
	const d = 3e20 / n
	fmt.Println(int64(n))
	// seperti var nilai bisa di dapatkan
	fmt.Println(math.Sin(n))
}
