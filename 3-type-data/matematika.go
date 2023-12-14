package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = a + b
	var d = a - b
	var e = a * b
	var f = a / b

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)

	i += 5 // i = i + 5
	fmt.Println(i)

	var j = i
	j++
	fmt.Println(j)
	j++
	fmt.Println(j)

	j--
	fmt.Println(j)
	j--
	fmt.Println(j)
}
