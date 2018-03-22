// https://tour.golang.org/basics/1

// Programs start running in package main.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator with the Unix time to ensure a different number every time.
	rand.Seed(time.Now().Unix())
	// The := syntax is shorthand for declaring and initializing a variable.
	// Go will infer the type of initialized variables.
	// This syntax can only be used inside a function.
	// var randIntn int = rand.Intn(100)
	randIntn := rand.Intn(100)

	fmt.Println("My favorite number is", randIntn)
}
