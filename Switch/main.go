package main

import (
	"fmt"
	"time"
)

func main() {

	// switch case teknik pertama
	// i=1 maka akan print case 1
	// i=2 maka akan print case 2
	// i=3 maka akan print case 3
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// switch case teknik kedua menentukan hari
	// case dengan enum
	// case bisa dengan beberapa nilai
	// jika case sebelumnya tidak memenuhi kondisi
	// maka akan print default
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// switch case teknik ketiga jam
	// case langsung pada variable
	// switch seperti ini baik di gunakan untuk handling error
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// switch case closure variable
	// switch case di olah dalam suatu variable closure
	// switch case menentukan tipe data dari paramater i
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm bool")
		case int:
			fmt.Println("I'm integer")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
