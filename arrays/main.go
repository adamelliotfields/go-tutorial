package main

import (
	"fmt"
)

// An array's length is part of its type, so it cannot be resized.
var words [2]string

// You can also declare and initialize an array literal.
var primes = [6]int{2, 3, 5, 7, 11, 13}

func main() {
	// You can assign array elements in local scope.
	words[0] = "Hello"
	words[1] = "World"

	// fmt.Println(words) // [Hello World]
	fmt.Println(words[0], words[1]) // Hello World

	// You can also declare and initialize an array using shorthand notation within local scope.
	// primes := [6]int{2, 3, 5, 7, 11, 13}

	fmt.Println(primes) // [2 3 5 7 11 13]
}
