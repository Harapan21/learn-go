package main

import "fmt"


// pointer
// & -> address 
// * -> pointer to address

func zeroVal(ival int) {
    ival = 0
}

func zeroPtr(iptr *int) {
    *iptr = 1
}

func main() {
    i := 2
    fmt.Println(i)

    zeroVal(i)
    fmt.Println(i) // -> masih tetap 2

    zeroPtr(&i)
    fmt.Println(i) // -> masih tetap 2

    z:= &i
}
