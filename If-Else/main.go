package main

import "fmt"

func main() {
	// jika 7 akar 2 sisa 0
	// maka akan print 7 is even 2
	// jika tidak  maka akan print 7 is odd
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// jika 8 akar 4 sisa 0 maka akan print is negative
	if 8%4 == 0 {
		fmt.Println("is negative")
	}

	// if else dengan deklarai langsung di kondisi
	// num sama dengan 9, jika num kurang 0
	// akan print num is negative
	// jika num kurawng dari 10 
	// akan print has 1 digit
	// jika semua tidak memenuhi kondisi maka 
	// akan print num has multiple digits
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	
}
