package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
		// Variables declared inside an `if` short statement are available inside any `else` blocks.
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	// Can't use v here, though.
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
