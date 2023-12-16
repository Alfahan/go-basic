package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	calculate := sumAll(10, 10, 10, 10, 10)
	fmt.Println(calculate)

	numbers := []int{10, 10, 10, 10}
	fmt.Println(sumAll(numbers...))
}
