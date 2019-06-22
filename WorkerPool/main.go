package main

import (
	"fmt"
	"time"
)

// menerima id for loop dari for 1,
// menerima jobs dari for 2,
// memasukan value id ke channel result
func worker(id int, jobs <-chan int, result chan<- int) {

	// for loop range jobs
	for j := range jobs {
		fmt.Println("worker", id, " started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		result <- j * 2
	}
}

func main() {
	// membuat channel maksimal number 100
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	// for 1
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, result)
	}
	// for 2
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// for 3
	for a := 1; a <= 5; a++ {
		<-result
	}
}
