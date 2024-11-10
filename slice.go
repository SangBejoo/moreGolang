package main

import "fmt"

func main() {
	// Slice ini merupakan potongan dari array
	// Slice mirip seperti array, namun ukurannya bisa berubah
	// Slice tidak perlu menentukan panjangnya
	// Slice selalu terkait dengan array
	// Slice selalu punya 3 data, yaitu pointer, length, dan capacity
	// Pointer adalah posisi awal elemen slice di array
	// Length adalah panjang dari slice
	// Capacity adalah panjang maksimal dari slice

	// Cara membuat slice dari array
	//array[low:high]
	names := [...]string{
		"A",
		"Y",
		"U",
		"B",
		"S",
		"U",
		"B",
		"A",
		"G",
		"I",
		"Y",
		"A",
	}
	slice := names[0:4]
	fmt.Println(slice)

	// deklarasi slice tanpa menggunakan array
	slice1 := names[4:12]
	fmt.Println(slice1)

	// deklarasi secara manual dari array yang sudah ada
	var slice2 []string = names[:]
	fmt.Println(slice2)

	// function len() untuk mengetahui panjang slice
	fmt.Println(len(slice))
	// function cap() untuk mengetahui kapasitas slice
	fmt.Println(cap(slice))
	// function append() untuk menambahkan elemen pada slice
	slice = append(slice, "Ayub")
	fmt.Println(slice2)
	fmt.Println("")
	//low adalah index awal
	//high adalah index akhir, namun tidak termasuk index akhir tersebut
	//jika low tidak ditulis maka slice akan dimulai dari index 0

	//kode program append slice
	days := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}
	daySlice1 := days[5:] //sabtu, minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Libur di sabtu"
	daySlice1[1] = "Libur di minggu"
	fmt.Println(daySlice1)

	daySlice2 := append(daySlice1, "Hari baru")
	daySlice2[0] = "sabtu lama"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)
	fmt.Println("")

	// make a slice
	var newSlice1 []string = make([]string, 2, 5)
	newSlice1[0] = "Ayyub"
	newSlice1[1] = "Subagiyya"
	newSlice2 := append(newSlice1, "Ayubsubagiya")
	fmt.Println(newSlice2)
	fmt.Println(newSlice1)
	fmt.Println(len(newSlice1))
	fmt.Println(cap(newSlice1))

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Ayub"
	newSlice[1] = "Subagiya"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2[0] = "budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice1)

	// copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// hati hati membuat slice
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
