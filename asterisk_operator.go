package main

import "fmt"

type Address1 struct {
	City, Province, Country string
}

func main() {
	address1 := Address1{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1 // copy value

	address2.City = "Bandung"
	fmt.Println(address1) // {Subang Jawa Barat Indonesia}
	fmt.Println(address2) // {Bandung Jawa Barat Indonesia}

	// pointer
	address3 := &address1
	address3.City = "Jakarta"

	fmt.Println(address1) // {Jakarta Jawa Barat Indonesia} ikut berubah
	fmt.Println(address3) // &{Jakarta Jawa Barat Indonesia}

	address2 = Address1{"Malang", "Jawa Timur", "Indonesia"}
	fmt.Println(address2) // &{Malang Jawa Timur Indonesia}

	*address3 = Address1{"Surabaya", "Jawa Timur", "Indonesia"}
	fmt.Println(address1) // {Surabaya Jawa Timur Indonesia}
	fmt.Println(address3) // &{Surabaya Jawa Timur Indonesia}
}
