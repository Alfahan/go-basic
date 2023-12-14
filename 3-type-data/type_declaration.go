package main

import "fmt"

func main() {
	type NoKtp string

	var ktpAlfa NoKtp = "111111111111"

	var contoh string = "222222222222"

	var contohKtp NoKtp = NoKtp(contoh)

	fmt.Println(ktpAlfa)
	fmt.Println(contohKtp)

}
