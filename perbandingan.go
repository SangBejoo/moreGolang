package main

import "fmt"

func main() {
	var name1 = "Ayub Subagiya"
	var name2 = "ayub subagiya"

	var result bool = name1 == name2
	fmt.Println(result)
	var result2 bool = name1 != name2
	fmt.Println(result2)
	var result3 bool = name1 < name2
	fmt.Println(result3)
	var result4 bool = name1 > name2
	fmt.Println(result4)
}
