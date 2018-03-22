package main

import (
	"fmt"
)

func main() {
	i := 42
	f := 3.142
	c := 0.867 + 0.5i

	fmt.Printf("%v is of type %T\n", i, i)
	fmt.Printf("%v is of type %T\n", f, f)
	fmt.Printf("%v is of type %T", c, c)
}
