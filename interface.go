package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHei(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name

}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Ayub"}
	sayHei(person)

	animal := Animal{Name: "Gajah"}
	sayHei(animal)
}
