package main

import "fmt"

func main() {
	name := "Ayub Sssubagiya"

	// kalau benar hanya if saja
	if name == "Ayub Subagiya" {
		fmt.Println("Hello Ayub Subagiya")
	} else { // kalau salah
		fmt.Println("Hello Stranger")
	}

	// Else if expression
	values := 90
	// kalau benar hanya if saja
	if values == 90 {
		fmt.Println("Nilai anda A")
	} else if values == 80 { // kalau salah
		fmt.Println("Nilai anda B")
	} else { // kalau salah
		fmt.Println("Nilai anda C")
	}

	// Else if expression
	names := "Subagiya"
	if names == "Ayubs" {
		fmt.Println("Hello Ayub Subagiya")
	} else if names == "Subagiya" {
		fmt.Println("Hello Subagiya")
	} else {
		fmt.Println("Hello Stranger")
	}

	// if dengan short statement
	namess := "Ayub Subagiya"
	if length := len(namess); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

	if length := len(names); length > 10 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
