package main

import (
	"fmt"
	"time"
)

func main() {
	// proses pertama
	c := make(chan string, 1) // membuat channel value string
	go func() {               // anonymose function
		time.Sleep(2 * time.Second) // delay time 2 second
		c <- "result 1"             // memasukan nilai "result 1" ke channel c
	}()

	select { // proses select
	case res := <-c: // menerima nilai channel c("result 1") ke res
		fmt.Println(res) // print nilai res
	case <-time.After(1 * time.Second): // jika delay waktu melebihi 1 second
		fmt.Println("timeout 1") // print timeout 1

	}

	// hasil proses pertama -> timeout 1 karena delay lebih dari 1 second
	// proses 2 sama seperti proses 1 hasil proses kedua -> result 2
	// karena delay tidak lebih dari 3 second
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
