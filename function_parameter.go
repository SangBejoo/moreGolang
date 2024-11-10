package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func numberAdd(number int, numberTwo int) {
	fmt.Println("Number", number+numberTwo)
}

func main() {
	sayHelloTo("Ayub", "Subagiya")
	numberAdd(10, 10)
}
