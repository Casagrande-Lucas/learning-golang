package main

import "fmt"

// Interfaces can hold values of any type that implements them.
// Type assertions and switches allow you to retrieve the underlying concrete type.

// Best Practice: Use type assertions and switches judiciously,
// primarily when dealing with interfaces that can hold multiple types.
// Favor type-safe designs when possible.

func describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case float64:
		fmt.Printf("Float: %.2f\n", v)
	default:
		fmt.Printf("Unknown type\n")
	}
}

func main() {
	describe(42)
	describe("hello")
	describe(3.14)
	describe(true)
}
