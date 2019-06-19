package main

import "fmt"

// fungsi pada go dengan format
// func NamaFunsi(param tipedata) tipedata nilai yang di return ( kembalikan ) {}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2", res)

	res := plusPlus(1, 2, 3)
	fmt.Println("1+2+3 = ", res)
}
