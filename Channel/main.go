package main

import (
	"fmt"
	"time"
)

// channel di istilahkan seperti pipa menghubungkan antara konkurensi ( goroutine )
// [goroutine] = channel = [goroutine]
// antara goroutine bisa mengoper nilai

func main() {
	// membuat channel
	messages := make(chan string)

	// goroutine <- ( memasukan ) string "ping" ke channel
	go func() { messages <- "ping" }()

	// variable := <- ( menerima ) value dari channel
	msg := <-messages
	fmt.Println(msg)

	// channel buffer
	messages2 := make(chan string, 2)
	messages2 <- "buffered"
	messages2 <- "channel"
	fmt.Println(<-messages2)
	fmt.Println(<-messages2)

	// Channel Synchronization
	done := make(chan bool, 1)
	go func(done chan bool) {
		fmt.Print("working...")
		time.Sleep(time.Second)
		fmt.Println("done")
		done <- true
	}(done)

	<-done

	// channel Direction
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	// param msg akan memasukan value string ke pings
	ping := func(pings chan<- string, msg string) {
		pings <- msg
	}
	// param pings menerima value dari channel bervalue string
	// variable msg menerima value dari param pings yang menerima value dari channel bervalue
	// string
	// msg memasukan value string ke pongs
	pong := func(pings <-chan string, pongs chan<- string) {
		msg := <-pings
		pongs <- msg
	}
	ping(pings, "passed message") // pings <- "passed message"
	pong(pings, pongs)            // pongs <- pings("passed message")
	fmt.Println(<-pongs)          // -> passed message
}
