package main

import "fmt"

func main() {
	// slice lebih powefull dari array
	// slice tidak sama seperti array
	// sesuai nama slice dapat memotong string utuh menjadi char
	s := make([]string, 3)
	fmt.Println("emp", s)

	// sama seperti untuk mengubah nilai
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
	fmt.Println(s[2])
	fmt.Println("len", len(s))

	// menambahkan nilai
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// copy value s ke c
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	// mengambil nilai s
	// dari index array 2
	// sampai index arrray 5 - 1 = 4
	l := s[2:5]
	fmt.Println(l)

	// memotong sampai ( tidak termasuk ) s[5]
	l = s[:5]
	fmt.Println(l)

	// memotong mulai dari ( termasuk  ) s[2]
	l = s[2:]
	fmt.Println(l)
	// deklarasi slice secara langsung
	t := []string{"g", "h", "t"}
	fmt.Println(t)


	
}
