package main

import "fmt"

func printAdd(x int, y int) {
	fmt.Println(x + y)
}

func add(x int, y int) int {
	return x + y
}

func addAndMultiply(x int, y int) (int, int) {
	return x + y, x * y
}

func main() {
	fmt.Println(add(42, 13))
	printAdd(42, 13)
	fmt.Println(addAndMultiply(42, 13))
}
