package main

import (
	"fmt"
)

// Numeric constants are high-precision values.
const (
	// Shift left 100
	Big = 1 << 100

	// Shift right 99
	Small = Big >> 99
)

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	// Result will be too big to print to stdout.
	// fmt.Println(needInt(Big))
	fmt.Println(needFloat(Big))
}

func needInt(x int) int {
	return (x * 10) + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}
