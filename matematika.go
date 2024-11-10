package main

import "fmt"

func main() {
	//Operator Aritmatika
	var a = 10
	var b = 10
	var c = a + b
	var d = (c - a) / b
	var e = c * d
	var f = e / c
	fmt.Println(f)
	fmt.Println(e)
	fmt.Println(d)
	fmt.Println(c)
	fmt.Println("")

	//Augmented Assignment adalah penambahan di variable sendiri
	var g = 19
	g += 10 // g = g + 10
	fmt.Println(g)

	var j = 20
	j -= 10 // j = j - 10
	fmt.Println(j)

	fmt.Println("")

	//unary operator
	var x = 10
	x++ // x = x + 1
	fmt.Println(x)
	x-- // x = x - 1
	fmt.Println(x)

}
