package main

import "fmt"

// Understanding pointers is vital for efficient memory usage and performance optimization.
// Pointers store memory addresses of variables.

func main() {
	var x int = 10
	var p *int = &x
	fmt.Println(*p) // Dereferencing: Outputs 10
	*p = 20
	fmt.Println(x) // Outputs 20
}
