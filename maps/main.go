package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

// A map maps keys to values.
// `make` returns an initialized map ready for use.
var vertexMap = make(map[string]Vertex)

// You can also make a map literal.
var vertexMapLiteral = map[string]Vertex{
	"google": {
		37.42202,
		-122.08408,
	},
}

func main() {
	// Use `map[key]` to insert or update an element in a map.
	vertexMap["bell_labs"] = Vertex{
		40.68433,
		-74.39967,
	}

	// You can also use `map[key]` to retrieve an element from a map.
	bellLabs := vertexMap["bell_labs"]
	google := vertexMapLiteral["google"]

	fmt.Printf("Bell Labs: %g,%g\n", bellLabs.Lat, bellLabs.Long) // Bell Labs: 40.68433,-74.39967
	fmt.Printf("Google: %g,%g\n", google.Lat, google.Long) // Google: 37.42202,-122.08408

	// Test that an element is present with a two-value assignment.
	// The first value is the element.
	// The second value is a boolean.
	if _, ok := vertexMap["bell_labs"]; ok {
		fmt.Println(ok) // true
	}

	// You can delete an element with `delete(map, key)`
	delete(vertexMap, "bell_labs")

	if _, ok := vertexMap["bell_labs"]; !ok {
		fmt.Println("key 'bell_labs' not found in map 'vertexMap'") // false
	}
}
