package main

import "fmt"



// menjalakan lebih dari satu fungsi secara berbarengan
func f(from string) {
  for i := 0; i<3; i++ {
    fmt.Println(from, ":", i)
  }
}

func main() {

      // secara lansung
      f ("direct")

      // menggunakan go routine
      go f("goroutine")

      // menggunakan goroutine anonymouse function
      go func(msg string) {
	fmt.Println(msg)
	
      }("goroutine2")


      fmt.Scanln()
      fmt.Println("done")
}
