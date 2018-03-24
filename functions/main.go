package main

import (
	"fmt"
)

// When consecutive parameters share a type, you only need to provide the type to the last one.
func add(x, y int) int {
	return x + y
}

// Functions can be passed to other functions.
func compute(fn func(x, y int) int) int {
	return fn(42, 42)
}

// Functions can return multiple results.
func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	// Return values can be "named".
	x = sum * 4 / 9
	y = sum - x

	// A return statement without arguments returns the named return values.
	return
}

func main() {
	sum := compute(add)

	fmt.Printf("The sum of 42 and 42 is %d\n", sum) // The sum of 42 and 42 is 84

	x, y := swap("world", "hello")

	fmt.Println(x, y) // hello world

	a, b := split(17)

	fmt.Println("a:", a) // a: 7
	fmt.Println("b:", b) // b: 10
}
