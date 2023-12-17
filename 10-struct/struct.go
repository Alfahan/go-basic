package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Helllo", name, "my name is", customer.Name)
}

func main() {
	var alfa Customer

	fmt.Println(alfa)

	alfa.Name = "Mohamad Ali Farhan"
	alfa.Address = "Indonesia"
	alfa.Age = 26

	fmt.Println(alfa)

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     29,
	}
	fmt.Println(joko)

	dyaksa := Customer{"Dyaksa", "Indonesia", 28}
	fmt.Println(dyaksa)

	dyaksa.sayHello("Alfa")
}
