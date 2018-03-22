package main

import (
	"fmt"
	"math"
)

func main() {
	i, j := 3, 4
	// The expression T(v) converts the value "v" to the type "T".
	f := math.Sqrt(float64((i * i) + (j * j)))
	u := uint(f)

	fmt.Println(i, j, f, u)
}
