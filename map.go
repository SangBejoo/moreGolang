package main

import "fmt"

func main() {
	//var person map[string]string = map[string]string{}
	//person["name"] = "Ayub Subagiya"
	//person["grade"] = "A"
	//person["status"] = "Graduated"
	//
	//fmt.Println()

	person := map[string]string{
		"name":   "Ayub Subagiya",
		"grade":  "A",
		"status": "Graduated",
	}
	fmt.Println(person["name"])
	fmt.Println(person["grade"])
	fmt.Println(person["status"])
	fmt.Println(person)

	fmt.Println(person["sasa"])

	// function len() digunakan untuk menghitung jumlah data pada map
	fmt.Println(len(person))
	// mengambil data di map dengan key
	fmt.Println(person["name"])
	// megubah data di map dengan key
	person["name"] = "Ayub"
	// membuat map baru dari map yang sudah ada
	var person2 = person
	person2["name"] = "Subagiya"
	fmt.Println(person2)
	fmt.Println(person)
	// menghapus data di map
	delete(person, "name")
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Ayub Subagiya"
	book["ups"] = "salah"

	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)

}
