package main

import (
	"fmt"
)

func main(){
	count, total := sum(1, 7, 3, 5, 9)
	fmt.Println(count, total)
}

func sum(nums ...int)(int, int){
	s := 0
	count := 0
	fot _, n := range nums {
		s += n
		count++
	}
	return count, s
}