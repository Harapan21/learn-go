package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	c1 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "two"
	}()

	// print value channel dengan value
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c:
			fmt.Println("received c: ", msg1)
		case msg2 := <-c1:
			fmt.Println("received c1: ", msg2)
		}
	}
}
