package main

import (
	"fmt"
	"net/url"
)

func main() {
	s := "https://gobyexample.com:8872/url-parsing"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.Hostname())
	fmt.Println(u.Port())
}
