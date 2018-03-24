package main

import (
	"fmt"
)

var power = make([]int, 10)

func main() {
	// The range form of the for loop iterates over a slice or map.
	for i := range power {
		power[i] = 1 << uint(i) // == 2**i
	}

	// `range` returns two values each iteration: the index and the value of the element.
	for i, v := range power {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// If you only want the value, use _ for the index.
	for _, value := range power {
		fmt.Printf("%d\n", value)
	}
}
