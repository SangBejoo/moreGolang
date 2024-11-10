package main

import "fmt"

func main() {
	type nama string

	var name nama = "1111111"
	fmt.Println(name)

	var contoh string = "A2121 Subagiya"
	var contoh1 nama = nama(contoh)
	fmt.Println(contoh1)
	fmt.Println(contoh)
}
