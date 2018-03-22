package main

import (
	"fmt"
)

// Constants can be char, string, bool, or int types.
// Constants use CamelCase names.
// Since upper-case identifiers are exported, use lower camelCase to make package-private.
const Pi = 3.14

func main() {
	// Constants cannot be declared using assignment shorthand syntax.
	const World = "World"
	const Truth = true

	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	fmt.Println("Go rules?", Truth)
}
