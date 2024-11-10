package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}
	fmt.Println("SELESAI")

	// for dengan statement post
	for angka := 1; angka <= 10; angka++ {
		fmt.Println("Perulangan Angka ", angka)
	}
	fmt.Println("SELESAI")

	// for range manual
	names := []string{"Ayub", "Subagiya", "Ayub Subagiya"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	fmt.Println("SELESAI")

	// for range
	names1 := []string{"Ayubsss", "subagiya", "Belajar Golang"}
	for i, names1 := range names1 {
		fmt.Println("Index ke ", i, "=", names1)
	}

	for _, names1 := range names1 {
		fmt.Println(names1)
	}

}
