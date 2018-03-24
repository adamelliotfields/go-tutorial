package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// Variable declarations can be factored into blocks, like import statements.
var (
	boolean bool       = false
	maxInt  uint64     = 1<<64 - 1
	complex complex128 = cmplx.Sqrt(-5 + 12i)
	i int // 0
	f float64 // 0
	b bool // false
	s string // ""
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", boolean, boolean)
	fmt.Printf("Type: %T Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v\n", complex, complex)

	a, b := 3, 4
	// The expression T(v) converts the value "v" to the type "T".
	c := math.Sqrt(float64((a * a) + (b * b)))
	d := uint(c)

	fmt.Println(a, b, c, d)

	x := 42
	y := 3.142
	z := 0.867 + 0.5i

	fmt.Printf("%v is of type %T\n", x, x)
	fmt.Printf("%v is of type %T\n", y, y)
	fmt.Printf("%v is of type %T\n", z, z)

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
