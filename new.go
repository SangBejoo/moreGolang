package main

import "fmt"

type Address12 struct {
	City, Province, Country string
}

func main() {
	var alamat1 *Address12 = new(Address12)
	var alamat2 *Address12 = alamat1

	alamat2.Country = "Indonesia"
	fmt.Println(alamat1) // &{ Indonesia}
	fmt.Println(alamat2) // &{ Indonesia}
}
