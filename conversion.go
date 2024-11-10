package main

import "fmt"

func main() {
	var nilai32 int32 = 32769
	var nilai64 int64 = int64(nilai32)
	var nilaiint16 int16 = int16(nilai32)
	var nilaiint8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilaiint16)
	fmt.Println(nilaiint8)

	var name = "Ayub Subagiya"
	var A = name[0]
	var aString = string(A)

	fmt.Println(name)
	fmt.Println(A)
	fmt.Println(aString)
}
