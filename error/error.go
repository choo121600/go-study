package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("1.txt")

	if err != nil {
		log.Fatal(err.Error())
	}
	println(f.Name())
	println("end")
}
