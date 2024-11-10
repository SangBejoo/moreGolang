package main

import "fmt"

func main() {
	const firstName string = "ayub"
	const lastName = "subagiya"

	const (
		age     = 20
		city    = "jakarta"
		address = "jl. abc"
	)
	fmt.Println(age)
	fmt.Println(city)
	fmt.Println(address)
}
