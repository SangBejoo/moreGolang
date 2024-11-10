package main

import "fmt"

func getGoodBye(name string) string {
	goodbye := "Goodbye " + name
	return goodbye
}

func main() {
	contoh := getGoodBye
	misal := getGoodBye

	fmt.Println(contoh("Ayub"))
	fmt.Println(misal("Subagiya"))

}
