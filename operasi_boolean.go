package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 65

	var lulusNilaiAkhir = nilaiAkhir > 75
	var lulusAbsensi = absensi > 75

	var lulus = lulusNilaiAkhir && lulusAbsensi
	fmt.Println(lulus)
}
