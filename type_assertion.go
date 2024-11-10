package main

import "fmt"

func random() any {
	return true
}

func main() {
	var result any = random()
	//var resultString string = result.(string)
	//fmt.Println(resultString)
	//
	//var resultInt int = result.(int)
	//fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")

	}

}
