package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Like `for` and `if`, you can run a short statement before the condition.
	// Go only runs the matching case - there is no fall-through.
	// This eliminates the need for a break statement.
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Go runs on OS X.")
	case "linux":
		fmt.Println("Go runs on Linux.")
	default:
		fmt.Printf("Go runs on %s.", os)
	}

	today := time.Now().Weekday()

	// `switch` evaluates cases from top to bottom.
	switch time.Saturday {
	case today + 0:
		fmt.Println("When's Saturday? Today!")
	case today + 1:
		fmt.Println("When's Saturday? Tomorrow.")
	case today + 2:
		fmt.Println("When's Saturday? In two days.")
	default:
		fmt.Println("When's Saturday? Too far away.")
	}

	now := time.Now()

	// `switch` without a condition is a clean way to write if-elseif-else chains.
	switch {
	case now.Hour() < 12:
		fmt.Println("Good morning!")
	case now.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
