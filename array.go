package main

import "fmt"

func main() {
	/*array diibaratkan sebagai lemari
	dimana setiap laci adalah elemen array
	dimulai dari 0
	*/
	//cara membuat array
	var names [5]string
	names[0] = "Ayub SUbagiya"
	names[1] = "Ayubs"
	names[2] = "Ayub"
	names[3] = "Subagiya"
	names[4] = "Ayub Subagiya"
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names[3])
	fmt.Println(names[4])

	//cara mengganti isi array
	names[2] = "aaaa"
	fmt.Println(names[2])

	//cara membuat array dengan langsung mengisinya
	var values = [3]int{
		10,
		20,
		30,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println("")
	//function len() untuk mengetahui panjang array
	fmt.Println(len(values))

	var penentuan = [...]int{
		10,
		20,
		30,
		40,
		50,
	}
	fmt.Println(penentuan)
	fmt.Println(penentuan[1])

}
