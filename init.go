package main

import (
	"UdemyGoland/database"
	_ "UdemyGoland/internal"
	"errors"
	"fmt"
)

func Pembagian(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	fmt.Println(database.GetDatabase())

}
