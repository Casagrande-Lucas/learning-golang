package main

import "fmt"

// a. Type Inference with :=

// Go's type system allows both explicit type declarations and type inference, enabling flexibility and safety.
// When using :=, Go infers the variable type based on the assigned value.

func main() {
	count := 10      // inferred as int
	price := 99.99   // inferred as float64
	name := "Gopher" // inferred as string
	isActive := true // inferred as bool

	fmt.Println(count, price, name, isActive)
}
