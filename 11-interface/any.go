package main

import "fmt"

func ups() any {
	//return 1
	return true
}

func main() {
	var kosong any = ups()
	fmt.Println(kosong)
}
