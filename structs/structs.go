package main

import (
	"fmt"
)

// A `struct` is a collection of fields.
type Person struct {
	Name string
	Age  int
}

func main() {
	adam := Person{"Adam Fields", 33}

	// Struct fields are accessed using dot notation.
	name := adam.Name

	// Struct fields can also be accessed through a pointer.
	p := &adam
	age := p.Age

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
}
