package main

import "fmt"

func main() {
	// tenik pengulangan (for) 1
	// explanation
	// deklarasi i dengan nilai 1
	// jika i kurang dari 3
	// print i lalu nilai i + 1
	// akan mengulangi sampai angka sama dengan 3
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// teknik pengulangan (for) 2
	// explaination
	// deklarasi j dengan nilai 7
	// j kurang dari 9
	// j++ ( j+1 )
	// akan mengulangi sampai angka sama dengan 9
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// teknik pengulangan (for) 3
	// explaination
	// mengulang tanpa logika
	// di akhiri dengan break untuk
	// menghentikan pengulangan

	for {
		fmt.Println("loop")
		break
	}

	//  teknik pengulangan (for) 4
	// sama dengan teknik pengulangan 3
	// tetapi di tambah proses logika if
	// jika nilai n akar 2 sisanya 0 maka lanjut berikutnya
	// tanpa menjalankan perintah selanjutnya
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
