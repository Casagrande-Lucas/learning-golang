package main

import "fmt"

// b. Explicit Typing

// Specifying the type can enhance clarity, especially in public APIs or when the inferred type isn't obvious.

// Best Practice: Use type inference for local variables with obvious types.
// Use explicit types for package-level variables, exported variables, or when clarity is needed.

var countExplicitType int = 10
var priceExplicitType float64 = 99.99
var nameExplicitType string = "Gopher"
var isActiveExplicitType bool = true

func main() {
	fmt.Println(countExplicitType, priceExplicitType, nameExplicitType, isActiveExplicitType)
}
