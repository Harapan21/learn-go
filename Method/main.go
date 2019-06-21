package main

import "fmt"


// Rect have width and height with type integer
type Rect struct {
  width, height int
}


// menerima data Rect dapat mengolah apa yang ada di dalam 
// struct rect yaitu width dan height
func (r *Rect) area() int {
  return r.width * r.height
}

func (r Rect) perim() int {
  return 2*r.width * 2*r.height
}

func main() {
  
    r := Rect{width:10, height:5}
    fmt.Println("area", r.area())	
    fmt.Println("perim", r.perim())


    rp := &r 
    fmt.Println("area", rp.area())
    fmt.Println("perim", rp.perim())
}

