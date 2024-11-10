package main

import (
	"UdemyGoland/helper"

	"fmt"
)

func main() {
	result := helper.SayGoodBye("Ayub")
	fmt.Println(result)

	result = helper.SayHello("Ayub")
	fmt.Println(result)

	result = helper.SayNoMore("Ayub")
	fmt.Println(result)

	fmt.Println(helper.Application)

}
