package main

import "fmt"

// pointer
// & -> address
// * -> pointer to address

func zeroVal(ival int) {
	// merubah nilai ival jadi 0
	ival = 0
	// ival berubah jadi 0 tetapi nilai i tidak berubah menjadi 0
}

// iptr -> (pointer ke address)
func zeroPtr(iptr *int) {
	// address <- 1
	*iptr = 1
}

// struct
type Vertex struct {
	x int
	y int
}

func main() {
	i := 2
	fmt.Println(i) // -> 2

	zeroVal(i)     // mengocopy nilai ke zeroVal
	fmt.Println(i) // -> masih tetap 2

	zeroPtr(&i)    // memasukan memory address ke zeroPtr
	fmt.Println(i) // -> 1

	v := Vertex{
		x: 20,
		y: 30,
	}
	p := &v
	p.x = 25
	fmt.Println(v)
}
