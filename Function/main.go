package main

import "fmt"

// fungsi pada go dengan format
// func NamaFunsi(param tipedata) tipedata nilai yang di return ( kembalikan ) {}
func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

// mengembalikan beberapa nilai
func vals() (int, int) {
	return 3, 7
}

// Variadic Functions
// ... artinya mengambil semua parameter
func sums(nums ...int) {
	fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
      total += num
    }
    fmt.Println(total)
}

// Closure function
func intSeq() func() int {
     i:= 0
     return func() int {
      i++
      return i
     }
}

// Recursion function
func fact(n int ) int {
  if n==0 {
    return 1
  }
  return n * fact(n-1)
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 = ", res)

	// parameter pertama mengambil nilai 3, paramater kedua mengambil nila 7
	_, c := vals()
	fmt.Println(c)

	nums := []int{1, 2, 3, 4, 5}
	// menaruh nums... menaruh [1,2,3,4,5] sebagai paramater si sums
	sums(nums...)

	// deklarasi function closure
	newInt := intSeq()
	fmt.Println(newInt())
}

