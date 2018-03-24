// Programs start running in package main.
package main

// You can group imports into a "factored" import statement.
// You can also write multiple import statements.
// import "fmt"
// import "math"
import (
	"fmt"
	"math"
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

	// A name is exported if it begins with a capital letter, like Pi.
	pi := math.Pi

	fmt.Printf("My favorite number is %d\n", randIntn)
	fmt.Printf("Pi is %g\n", pi)
	fmt.Printf("The square root of 7 is %g\n", math.Sqrt(7))
}
