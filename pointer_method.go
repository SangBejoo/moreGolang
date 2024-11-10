package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	ayub := Man{"Ayub"}
	ayub.Married()

	fmt.Println(ayub.Name) // Ayub
}
