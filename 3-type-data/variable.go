package main

import "fmt"

func main() {
	//opsi 1
	var name string
	name = "Mohamad Ali"
	fmt.Println(name)

	name = "Ali Farhan"
	fmt.Println(name)

	//opsi 2
	var name2 = "Moh Ali"
	fmt.Println(name2)

	name2 = "Ali Farhan"
	fmt.Println(name2)

	//opsi 3
	name3 := "M Ali"
	fmt.Println(name3)

	name3 = "Ali Farhan"
	fmt.Println(name3)

	//opsi 4
	var (
		firstName  = "Mohamad"
		middleName = "Ali"
		lastName   = "Farhan"
	)

	fmt.Println(firstName + " " + middleName + " " + lastName)
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

}
