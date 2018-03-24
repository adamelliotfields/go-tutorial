package main

import (
	"fmt"
)

// The var statement declares a list of variales and can be at the package or function level.
// The type definition is last.
var c, python, java bool

var i int

// If a variable is initialized, the type can be omitted.
var a, b = 1, 2

func main() {
	fmt.Println(i, a, b) // 0 1 2

	// Variables can be assigned inside a function block.
	c, python = true, false

	// Inside a function, the short assignment statement can be used instead of var.
	x := 1
	y, z := 2, 3
	java := "NO!"

	fmt.Println(x, y, z, c, python, java) // 1 2 3 true false NO!
}
