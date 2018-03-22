package main

import (
	"fmt"
)

// If a variable is initialized, the type can be omitted.
var i, j = 1, 2

func main() {
	var c, python, java = true, false, "no!"

	fmt.Println(i, j, c, python, java)
}
