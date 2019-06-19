package main

import "fmt"

// struct
// biasanya untuk struktur data
type Person struct{
 name string
 age int
}

func main() {
  fmt.Println(Person{"bob", 27})

  fmt.Println(Person{name:"Alice", age:30})

  s:= Person{name:"Sean", age: 50}
  fmt.Println(s)


  sp := &s
  fmt.Println(sp)

  sp.age = 51
  fmt.Println(sp)
}
