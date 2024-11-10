package main

import "fmt"

// nill hanya bisa di gunakan pada tipe data interface, function, map, slice, pointer, dan channel
// nill artinya kosong

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("")

	if data == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(data["name"])
	}
}
