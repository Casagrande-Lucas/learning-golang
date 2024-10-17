package main

import "fmt"

// a. Implicit vs. Explicit Conversion

// Go is a statically typed language with strict type safety, preventing unintended type mismatches.
// However, type conversions are necessary when working with different types.

// Go does not perform implicit type conversions, all conversions must be explicit.

func main() {
	var i int = 42
	var f float64

	// f = i // Error: cannot use i (type int) as type float64
	f = float64(i) // Correct

	fmt.Println(f)
}
