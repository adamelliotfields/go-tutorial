package main

import (
	"fmt"
)

// A `struct` is a collection of fields.
type Person struct {
	Name string
	Age  int
}

// If you know the order of the struct fields, you don't have to declare them.
// var adam = Person{Name: "Adam Fields", Age: 33}
var adam = Person{"Adam Fields", 33}

// Struct fields are accessed using dot notation.
var name = adam.Name

// Struct fields can also be accessed through a pointer.
var adamPointer = &adam
var age = adamPointer.Age

// Struct fields can also be functions.
type Counter struct {
	value     func() int
	increment func() int
	decrement func() int
}

// Returns a Struct with 3 methods.
func counter() Counter {
	v := 0

	return Counter{
		value:     func() int { return v },
		increment: func() int { v += 1; return v },
		decrement: func() int { v -= 1; return v },
	}
}

var c = counter()

func main() {
	fmt.Println("Name:", name) // Name: Adam Fields
	fmt.Println("Age:", age) // Age: 33

	fmt.Println("Counter:")

	for i := 0; i < 3; i++ {
		c.increment()
		fmt.Println(c.value())
	}
}
