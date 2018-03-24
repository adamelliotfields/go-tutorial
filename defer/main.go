package main

import (
	"fmt"
)

func deferredHello() {
	// Defers the execution until the surrounding function returns.
	defer fmt.Println("World")

	fmt.Println("Hello")
}

func deferredCounter() {
	fmt.Println("Counting...")

	for i := 1; i <= 10; i++ {
		// Deferred calls are pushed onto a stack and executed in last-in-first-out order.
		defer fmt.Println(i) // 10 9 8 7 6 5 4 3 2 1
	}

	fmt.Println("Done!")
}

func main() {
  deferredHello()
  deferredCounter()
}
