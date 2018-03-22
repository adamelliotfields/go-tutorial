package main

import (
	"fmt"
)

func main() {
	fmt.Println("Counting...")

	for i := 1; i <= 10; i++ {
		// Deferred calls are pushed onto a stack and executed in last-in-first-out order.
		defer fmt.Println(i)
	}

	fmt.Println("Done!")
}
