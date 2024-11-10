package main

import "fmt"

func getFullName() (string, string) {
	return "Ayub", "Subagiya"
}

func getNama() (string, string) {
	return "Ayub", "Subagiya"
}

func nilaiKelas() (int, int) {
	return 100, 90
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	namaDepan, namaBelakang := getNama()
	fmt.Println(namaDepan, namaBelakang)

	firstName, _ = getFullName()
	fmt.Println(firstName)

	_, lastName = getFullName()
	fmt.Println(lastName)

	_, nilai := nilaiKelas()
	fmt.Println(nilai)

	nilai, _ = nilaiKelas()
	fmt.Println(nilai)
}
