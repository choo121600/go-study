package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5)
	printSlice("a", a)
}

func printSlice(s string, x []int) {
	fmt.Println("%s len=%d cap=%d %v\n")
}
