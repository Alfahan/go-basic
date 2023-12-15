package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	fmt.Println("Selesai")

	for coba := 1; coba <= 10; coba++ {
		fmt.Println("Perulangan coba ke ", coba)
	}
	fmt.Println("Selesai")

	names := []string{"Mohamad", "Ali", "Farhan"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
