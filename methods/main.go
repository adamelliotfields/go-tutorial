package main

import (
	"fmt"
)

// Go does not have classes, so there are no instance methods.
// However, you can define methods on types.
type Greeting struct {
	Message string
}

// A method is a function with a "receiver" argument.
// The receiver appears between the `func` keyword and method name.
// In this example, `(g Greeting)` is the receiver.
// This is known as a "value" receiver.
func (g Greeting) Greet() {
	fmt.Println(g.Message)
}

// A "pointer receiver" acts like a setter method in traditional object-oriented languages.
// Pointer receivers are recommended, even as getters, as they don't need to copy the value on each call.
// It is recommended to only use one type of receiver on a given type (either value or pointer),
// because an interface cannot define mixed receivers.
func (g *Greeting) ChangeGreeting(newMessage string) {
	g.Message = newMessage
}

// Methods can be declared on non-structs as well.
type Name string

// Don't use a pointer here, as `n` will reference the memory address, not the value (the string) if you do.
func (n Name) Print() {
	fmt.Println(n)
}

var g = Greeting{"Hello World"}
var n = Name("Adam")

func main() {
	g.Greet() // Hello World

	// As a convenience, you don't need to use the pointer operator (&) when invoking the pointer receiver function.
	// (&g).ChangeGreeting("Good Bye World")
	g.ChangeGreeting("Good Bye World")

	g.Greet() // Good Bye World

	n.Print() // Adam
}
