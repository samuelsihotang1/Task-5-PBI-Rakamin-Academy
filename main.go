package main

import (
	"belajar-go/initializers"
	"fmt"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	fmt.Println("Hello, World!2")
}