package main

import "fmt"

func main() {
	x := 2
	perulangan(x)
}

func nambah(x int) int {
	return x + 1
}

func perulangan(i int) {
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
}
