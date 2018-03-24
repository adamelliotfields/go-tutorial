package main

import (
	"fmt"
	"strings"
)

var names = [4]string{"John", "Paul", "George", "Ringo"}

func main() {
	// Unlike arrays, slices can be dynamically sized.
	// A slice does not store any data, it just describes a section of an underlying array.
	namesSlice := names[0:1] // [John]

	// You can also use slice defaults to omit the low or high bounds of the slice.
	// namesSlice := namesArray[:1] // [John]

	// Changing the elements of a slice modifies the elements of its underlying array.
	namesSlice[0] = "Beatle"

	fmt.Println(namesSlice) // [Beatle]

	// A slice literal is like an array literal without the length.
	// Slice literals create the underlying array and then build a slice of that array.
	// Note: this is a really ugly way to declare a struct literal - use a type if possible.
	structSlice := []struct {
		key   int
		value bool
	}{
		{0, false},
		{1, true},
	}

	// You can also slice a slice.
	structSliceSlice := structSlice[:1]

	fmt.Println(structSliceSlice[0]) // {0, false}

	// Slices have a length and a capacity.
	// The length of a slice is the number of elements it contains.
	// The capacity is the number of elements in the underlying array,
	// counting from the low-bound of the slice.
	fmt.Println("Length:", len(namesSlice))   // Length: 1
	fmt.Println("Capacity:", cap(namesSlice)) // Capacity: 4

	// The zero-value of a slice is `nil`.
	// A nil slice has a length and capacity of 0.
	var nilSlice []int

	if nilSlice == nil {
		fmt.Println("nil") // nil
	}

	fmt.Println(len(nilSlice), cap(nilSlice)) // 0 0

	// make() takes the type, length, and capacity as arguments.
	// This creates a zeroed-array and then a slice of that array with a length of 5.
	makeSlice := make([]int, 5, 7)

	fmt.Println(makeSlice) // [0 0 0 0 0]

	// You can also make slices of slice literals.
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// Players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := range board {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// The return value of append is a slice containing all the elements of the original slice,
	// plus the new values.
	var newSlice []int
	newSlice = append(newSlice, 1, 2, 3)

	fmt.Println(newSlice) // [1 2 3]
}
