package main

import "fmt"

func main() {
	// array
	var names [3]string

	names[0] = "Mohamad"
	names[1] = "Ali"
	names[2] = "Farhan"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{
		1,
		2,
		4,
		3,
		5,
	}

	fmt.Println(values)
	fmt.Println(len(values))
}
