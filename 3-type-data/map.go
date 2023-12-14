package main

import "fmt"

func main() {
	//var person map[string]string = map[string]string{}
	//person["name"] = "Alfa"
	//person["address"] = "Subang"

	person := map[string]string{
		"name":    "Alfa",
		"address": "Subang",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Alfa"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}
