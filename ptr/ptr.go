package main

import (
	"fmt"
)

func main() {
	var num int = 1
	var numPtr *int = &num

	fmt.Println(numPtr)
	fmt.Println(*numPtr)
	fmt.Println(&num)
}
