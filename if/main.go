package main

import (
	"fmt"
	"math"
)

func recursiveSqrt(x float64) string {
	if x < 0 {
		return recursiveSqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// `if` statements can start with a short statement to execute before the condition.
	if v := math.Pow(x, n); v < lim {
		return v
	}

	return lim
}

func powWithElse(x, n, lim float64) float64 {
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
	fmt.Println(recursiveSqrt(2), recursiveSqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println(
		powWithElse(3, 2, 10),
		powWithElse(3, 3, 20),
	)
}
