package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	generate()
}

//go:generate go run main.go generate

func generate() {
	fmt.Println("Generated code")
}
