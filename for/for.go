package main

import (
	"fmt"
)

func main() {
	sum := 1
	cnt := 0
	for sum < 1000 {
		sum += sum
		cnt++
	}
	fmt.Println(sum)
	fmt.Println(cnt)
}
