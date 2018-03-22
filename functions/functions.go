package main

import (
	"fmt"
)

func main() {
	sum := add(42, 42)

	fmt.Println("The sum of 42 and 42 is", sum)
}

// When consecutive parameters share a type, you only need to provide the type to the last one.
func add(x, y int) int {
	return x + y
}
