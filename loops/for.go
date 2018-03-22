package main

import (
	"fmt"
)

func main() {
	// Go has only one looping construct - the `for` loop.
	// The basic `for` loop has three components:
	//   Init: executed before the first iteration
	//   Condition: executed before every iteration
	//   Post: executed at the end of every iteration
	// The loop ends when the condition evaluates to false.
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}
