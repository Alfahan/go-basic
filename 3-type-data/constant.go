package main

import "fmt"

func main() {
	const firstName string = "Ali"
	const lastName = "Farhan"

	fmt.Println(firstName)
	fmt.Println(lastName)

	const (
		firstName1 string = "Ali"
		lastName1  string = "Farhan"
	)

	fmt.Println(firstName1)
	fmt.Println(lastName1)
}
