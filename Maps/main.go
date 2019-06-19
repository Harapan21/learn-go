package main

import "fmt"
func main() {
	// map seperti array tapi memiliki kunci / id pada value
	// map[ "kunci":"value" ]
	m := make(map[string]int)

	// menambangkan kunci:value pada map
	// m[kunci] = value
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println(m)
	
	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len:", len(m))

	// mengapus value di m dengan kunci k2
	delete(m, "k2")
	fmt.Println("map: ", m)

	// cek value k2 ada atau tidak 
	// hasilnya bool (true | false)
	_,prs := m["k2"]
	fmt.Println("prs:", prs)

	/// membuat map secara langsung
	// map[tipedata kunci]tipedata value {kunci:value}
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ",n)
}
