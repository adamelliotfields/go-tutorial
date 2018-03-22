package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	// A `for` loop without any conditions will loop forever.
	for {
		time.Sleep(time.Second)
		fmt.Println(i)
		i += 1
	}
}
