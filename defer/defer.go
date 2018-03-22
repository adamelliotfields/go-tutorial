package main

import (
	"fmt"
)

func main() {
	// Defers the execution until the surrounding function returns.
	defer fmt.Println("World")

	fmt.Println("Hello")
}
