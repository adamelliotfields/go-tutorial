package main

import (
	"fmt"
	"math"
)

func main() {
	// A name is exported if it begins with a capital letter, like Pi.
	pi := math.Pi

	fmt.Println("Pi is", pi)
}
