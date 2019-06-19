package main

import "fmt"

func main() {
	// range
	// fungsi interator pada go
	nums := []int{1, 2, 3}
	sum := 0

	// range manghasilkan index,value
	// setiap item terdapat pada array
	for _, num := range nums {
		sum += num
	}

	fmt.Println(sum)

	kvs := map[string]string{"a": "Apple", "b": "Banana"}

	// pada slice akan menerima kunci, value
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	// k sebagai kunci
	for k := range kvs {
		fmt.Println("key:", k)
	}
	for i, c :=range "go"{
	  fmt.Println("key:", i,c )
	}

}
