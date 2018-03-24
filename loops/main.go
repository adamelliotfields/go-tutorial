package main

import (
	"fmt"
	"time"
)

func main() {
	// Go has only one looping construct - the `for` loop.
	// The basic `for` loop has three components:
	//   - Init: executed before the first iteration
	//   - Condition: executed before every iteration
	//   - Post: executed at the end of every iteration
	// The loop ends when the condition evaluates to false.
	fmt.Println("FIZZBUZZ:")
	for i := 1; i <= 30; i++ {
		fizzBuzz := ""

		if i%3 == 0 { fizzBuzz += "Fizz" }
		if i%5 == 0 { fizzBuzz += "Buzz" }

		if fizzBuzz != "" {
			fmt.Println(fizzBuzz)

			// `continue` exits the loop and starts the next iteration.
			// `break` exits the loop, but continues function execution.
			// `return` exits the loop, and stops functions execution.
			continue
		}

		fmt.Println(i)
	}

	// Go has no formal `while` loop, but a `for` loop can be used as such.
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Printf("\nWHILE LOOP:\n")
	fmt.Println(sum)

	// A `for` loop without any conditions will loop forever.
	i := 1
	fmt.Printf("\nINFINITE LOOP:\n")
	for {
		time.Sleep(time.Second)
		fmt.Println(i)
		i += 1
	}
}
