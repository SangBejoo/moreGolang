package helper //harus sama dengan folder name

// harus menggunakan kapital untuk diakses di package lain
var version = "1.0.0"

// ini benar
var Application = "Golang"

func SayHello(name string) string {
	return "Hello " + name
}

func SayGoodBye(name string) string {
	return "Goodbye " + name
}

func SayNoMore(name string) string {
	return "No More " + name
}
