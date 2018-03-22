package main

import (
	"fmt"
)

// The var statement declares a list of variales and can be at the package or function level.
// The type definition is last.
var c, python, java bool

func main() {
	var i int

	// Defaults to 0 false false false.
	fmt.Println(i, c, python, java)
}
