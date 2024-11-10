package main

func getCompleteName() (firstName, middleName, lastName string, age int) {
	firstName = "Ayub"
	middleName = "Subagiya"
	lastName = "Umur saya"
	age = 20
	return firstName, middleName, lastName, age
}

func main() {
	a, b, c, d := getCompleteName()
	println(a, b, c, d)
}
