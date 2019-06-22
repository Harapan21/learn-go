package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer 1 expired")
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	isStop2 := timer2.Stop()
	if isStop2 {
		fmt.Println("Timer 2 Stop")
	}
}
