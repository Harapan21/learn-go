package main

import "fmt" 
func main() {
	// slice lebih powefull dari array
	// slice tidak sama seperti array
	// setiap element yang ada di slice harus 1 tipe data
	// sesuai nama slice dapat memotong string utuh menjadi char
	s := make([]string, 3)
	fmt.Println("emp", s)


	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
}
