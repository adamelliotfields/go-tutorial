package main

import (
	"fmt"
)

func main() {
	a, b := split(17)

	fmt.Println("a:", a)
	fmt.Println("b:", b)
}

func split(sum int) (x, y int) {
	// Return values can be "named".
	x = sum * 4 / 9
	y = sum - x

	// A return statement without arguments returns the named return values.
	return
}
