package main

import (
	"fmt"
)

// Constants can be char, string, bool, or int types.
// Constants use CamelCase names.
// Since upper-case identifiers are exported, use lower camelCase to make package-private.
const Pi = 3.14

// Numeric constants are high-precision values.
const (
	// Shift left 100
	Big = 1 << 100

	// Shift right 99
	Small = Big >> 99
)

func needInt(x int) int {
	return (x * 10) + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// Constants cannot be declared using assignment shorthand syntax, even in local scope.
	const World = "World"
	const Truth = true

	fmt.Printf("Hello %s\n", World)
	fmt.Printf("Happy %g Day\n", Pi)
	fmt.Printf("Go rules? %t\n", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	// Result will be too big to print to stdout.
	// fmt.Println(needInt(Big))
	fmt.Println(needFloat(Big))
}
