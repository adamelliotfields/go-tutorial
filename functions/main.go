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

// Go functions can be closures as well.
func fibonacci(f1, f2 int) func() int {
	var first, second, next int

	// Fibonacci can be seeded with either 0,1 or 1,1
	first = f1
	second = f2

	return func() int {
		next = first
		first = second
		second = first + next

		return next
	}
}

func main() {
	sum := compute(add)
	fmt.Printf("The sum of 42 and 42 is %d\n", sum) // The sum of 42 and 42 is 84

	x, y := swap("world", "hello")
	fmt.Println(x, y) // hello world

	a, b := split(17)
	fmt.Printf("17 split is %d and %d\n", a, b) // 17 split is 7 and 10

	f := fibonacci(0, 1)
	fmt.Println("Fibonacci:")
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
