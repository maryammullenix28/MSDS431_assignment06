//go:build ignore
// +build ignore

package main

import (
	"fmt"
)

//go:generate go run main.go generate

func GenerateCode() {
	fmt.Println("Generated code")
}
