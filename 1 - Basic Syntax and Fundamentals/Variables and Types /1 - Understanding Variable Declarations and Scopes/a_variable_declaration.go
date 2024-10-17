package main

import "fmt"

// a. Variable Declaration

// In Go, variables can be declared using the var keyword or the short declaration :=.
// Understanding the nuances between these methods is crucial for writing clear and efficient code.

// Using var
var ageAlice int = 30

// Type inference with var
var nameAlice string = "Alice"

func main() {
	// Short declaration (only inside functions)
	heightAlice := 175.5

	fmt.Println(heightAlice, nameAlice, ageAlice)
}
