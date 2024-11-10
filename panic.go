package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Message:", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("APP ERROR")
	}

}

func main() {
	runApp(true)

}
