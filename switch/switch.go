package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")

	// Like `for` and `if`, you can run a short statement before the condition.
	switch os := runtime.GOOS; os {
	// Go only runs the matching case - there is no fall-through.
	// This eliminates the need for a break statement.
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
}
