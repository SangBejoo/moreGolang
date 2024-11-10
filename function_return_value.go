package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	result := getHello("Ayubs")
	fmt.Println(result)

	fmt.Println(getHello("Subagiya"))
	fmt.Println(getHello("Budi"))
}
