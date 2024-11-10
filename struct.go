package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	Married       bool
	Pekerjaan     string
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	var ayub Customer

	fmt.Println(ayub)

	ayub.Name = "Ayub Subagiya"
	ayub.Age = 20
	ayub.Address = "Indonesia"
	fmt.Println(ayub.Name)
	fmt.Println(ayub.Address)

	subagiya := Customer{
		Name:    "subagiya",
		Address: "Indonesia",
		Age:     30,
	}
	fmt.Println(subagiya)

	ayub1 := Customer{"subagiya", "Indonesia", 30, false, "Intern"}
	fmt.Println(ayub1)

	ayub2 := Customer{
		Name:    "subagiya",
		Address: "Indonesia",
		Age:     30,
	}
	fmt.Println(ayub2)

	// struct method
	ayub.sayHi("Aguss")
	ayub2.sayHi("Aguss")
	ayub1.sayHi("Aguss")

}
