package main

import (
	"fmt"
)

func main() {
	a, b := swap("world", "hello")

	fmt.Println(a, b)
}

// Functions can return multiple results.
func swap(x, y string) (string, string) {
	return y, x
}
