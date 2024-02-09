package main

import (
	"fmt"
	"go-basic/15-import/helper"
)

func main() {
	result := helper.SayHello("Alfa")
	fmt.Println(result)

	fmt.Println(helper.Application)
	fmt.Println(helper.vesion)            // tidak bisa diakses
	fmt.Println(helper.sayGoodBye("Eko")) // tidak bisa diakses
}
