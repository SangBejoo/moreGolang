package main

import "fmt"

func main() {
	name := "Ayub Subagiya"

	switch name {
	case "Ayub Subagiya":
		fmt.Println("Hello Ayub Subagiya")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Kenalan dulu?")
	}

	values := 90

	switch values {
	case 90:
		fmt.Println("Perfect")
	case 80:
		fmt.Println("Good")
	default:
		fmt.Println("Not Bad")

	}

	names := "Subagiya"
	switch names {
	case "Ayub":
		fmt.Println("Hello Ayub")
	case "Subagiya":
		fmt.Println("Hello Subagiya")
	default:
		fmt.Println("Kenalan dulu?")

	}

	// switch dengan short statement
	switch length := len(names); length > 5 {
	case true:
		fmt.Println("Terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	default:
		fmt.Println("Kenalan dulu?")
	}

	// switch tanpa kondisi
	names = "Subagiya ayyub"
	length := len(names)
	switch {
	case length > 10:
		fmt.Println("Terlalu panjang")
	case length > 5:
		fmt.Println("Nama sudah benar")
	default:
		fmt.Println("Kenalan dulu?")
	}
}
