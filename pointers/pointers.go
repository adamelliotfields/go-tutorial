package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	// The `&` operator generates a pointer to its operand.
	p := &i

	// The `*` operator denotes the pointer's underlying value.
	fmt.Println(*p)

	// You can change the value of `i` via the pointer.
	*p = 21
	fmt.Println(i)

	// Change `p` to point to `j`
	p = &j

	// Change `j` by dividing itself by 37 via the pointer.
	*p = *p / 37
	fmt.Println(j)
}
