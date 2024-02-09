package main

import (
	"fmt"
	"go-basic/16-package-initialization/database"
	_ "go-basic/16-package-initialization/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}
