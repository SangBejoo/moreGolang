package main

import "fmt"

type Address123 struct {
	City, Province, Country string
}

func changeCountryToIndonesia(address *Address123) {
	address.Country = "Indonesia"
}

func main() {
	var address Address123 = Address123{}
	changeCountryToIndonesia(&address)
	fmt.Println(address) // {Subang Jawa Barat Indonesia}
}
