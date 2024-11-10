package main

import "fmt"

func sumAll(number ...int) int {
	total := 0
	for _, num := range number {
		total += num
	}
	return total
}

func fullName(name ...string) string {
	fullName := ""
	for _, n := range name {
		fullName += n + " "
	}
	return fullName
}

func main() {
	fmt.Println(sumAll(10, 10, 10, 10, 10))
	fmt.Println(sumAll(10, 10, 10))
	fmt.Println(sumAll(10, 10))
	fmt.Println(sumAll(10))

	fmt.Println(fullName("Ayub", "Subagiya"))
	fmt.Println(fullName("Ayub", "Subagiya", "Belajar", "Golang"))
	fmt.Println(fullName("Ayub", "Subagiya", "Belajar", "Golang", "di", "Golang"))

	// slice parameter
	numbers := []int{10, 10, 10, 10, 10}
	fmt.Println(sumAll(numbers...))

}
