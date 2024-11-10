package main

import "fmt"

func main() {
	// break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("SELESAI")
	// continue
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("ganjil ke", i)
	}
	fmt.Println("SELESAI")

	// continue and break
	for i := 0; i <= 20; i++ {
		if i%2 == 0 {
			continue
		} else if i == 15 {
			break
		}
		fmt.Println("ganjil ke", i)

	}
}
