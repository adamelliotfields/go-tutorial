package main

import (
	"fmt"
)

func main() {
	sum := 1

	// Go has no formal `while` loop, but a `for` loop can be used as such.
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}
