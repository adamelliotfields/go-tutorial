package main

import (
	"fmt"
)

func main() {
	// Inside a function, the short assignment statement can be used instead of var.
	i := 1
	j, k := 2, 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
