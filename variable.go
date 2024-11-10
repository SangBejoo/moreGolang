package main

import "fmt"

func main() {
	var name = "Ayub Subagiya"
	fmt.Println(name)

	name = "Ayub"
	fmt.Println(name)

	var price = 10
	fmt.Println(price)

	prices := 20
	fmt.Println(prices)

	prices = 30
	fmt.Println(prices)

	var (
		firstName  = "Ayub"
		middlename = "And"
		lastName   = "Subagiya"
	)
	fmt.Println(firstName)
	fmt.Println(middlename)
	fmt.Println(lastName)
}
