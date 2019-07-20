package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// go func(messages chan<- string) {
	// 	messages <- "this is messae for messages channel"
	// }(messages)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg1 := "hi"

	select {
	case messages <- msg1:
		fmt.Println("sent message", <-messages)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
