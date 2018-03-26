package main

import (
	"encoding/json"
	"fmt"
)

// An interface type is a set of method signatures.
// Any type that defines all of the methods in the interface will implicitly implement the interface.
// Unlike other object-oriented languages, there is no "implements" keyword in Go.
// A good naming convention for interfaces is to end in "-er", since interfaces describe actions.
// In this example, anything that implements the JSONPrinter interfaces is a "JSON Printer".
type JSONPrinter interface {
	toJSON() string
}

// The field names must be capitalized so they can be exported.
// Use a "field tag" to describe how the generated JSON should be formatted.
// This is for the Marshal function from the json package, not specific to interfaces in general.
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) toJSON() string {
	// Marshal returns a []byte or error
	b, err := json.Marshal(u)

	if err != nil {
		// The Error() method returns a string of the error
		return err.Error()
	}

	return string(b)
}

// Instead of passing the User type to the printJSON function, use the JSONPrinter interface.
// This means that printJSON can accept any type that implements the JSONPrinter interface,
//   (has a toJSON() method).
func printJSON(j JSONPrinter) {
	fmt.Println(j.toJSON())
}

// Interface values can be thought of as a tuple of (value, type).
// This function prints the value and type of the underlying type implementing the interface.
// Note the use of the empty interface type, which can be used as a wildcard type,
//   like <?> in Java or any in TypeScript.
// This works because every type in Go implements at least zero methods.
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	// A `nil` interface holds neither value nor concrete type.
	var j JSONPrinter
	var u *User

	describe(j) // (<nil>, <nil>)

	j = u
	describe(j) // (<nil>, *main.User)

	j = &User{"adam", "gopher"}
	describe(j) // (&{adam gopher}, *main.User)

	printJSON(j) // {"username":"adam","password":"gopher"}
}
