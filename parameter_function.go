package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "anjing" || name == "babi" || name == "bangsat" || name == "kontol" {
		return "..."
	} else {
		return name
	}
}

func getName(names string, filter1 Filter) {
	filteredName := filter1(names)
	fmt.Println("Hello", filteredName)
}

func main() {
	sayHelloWithFilter("anjing", spamFilter)
	sayHelloWithFilter("Ayub", spamFilter)
	spamFilter("babi")

	filter := spamFilter
	sayHelloWithFilter("babi", filter)
	sayHelloWithFilter("ayub subagiya", filter)

	dapetName := getName
	dapetName("ayub subagiya", spamFilter)

	//type declaration

}
