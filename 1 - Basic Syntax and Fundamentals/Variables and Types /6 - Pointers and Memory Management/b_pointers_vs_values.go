package main

import "fmt"

// Deciding whether to pass by pointer or by value affects performance and behavior.

// Use Cases:
// • Large Structs: Pass pointers to avoid copying.
// • Mutating Data: Use pointers when the function needs to modify the original data.
// • nil Pointers: Represent absence of value or optional data.

// Best Practice: Use pointers for large data structures and when you need to modify the original data.
// Use values for small, immutable data.

type Data struct {
	Values [1000]int
}

func processByValue(d Data) {
	fmt.Println("processByValue: Inefficient for large structs. Copies entire Data struct: ", d.Values[0])
}

func processByPointer(d *Data) {
	// Operates on original Data struct
	fmt.Println("processByPointer: Efficient and allows mutation. Operates on original Data struct: ", d.Values[0])
}

func main() {
	d := Data{}
	processByValue(d)    // Inefficient for large structs
	processByPointer(&d) // Efficient and allows mutation
}
