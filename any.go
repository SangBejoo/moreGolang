package main

import "fmt"

func Ups() any {
	return "ups"
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}
