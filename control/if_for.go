package main

import (
	"fmt"
)

func printhello(x int) int {
	for i := 0; i < x; i++ {
		if i%2 == 0 {
			fmt.Println("Hello World")
		}
	}
	return x
}

func main() {
	printhello(5)
}
