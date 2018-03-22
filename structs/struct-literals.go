package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Age will default to 0
	adam := Person{Name: "Adam Fields"}

	fmt.Println("Name:", adam.Name)
	fmt.Println("Age:", adam.Age)
}
