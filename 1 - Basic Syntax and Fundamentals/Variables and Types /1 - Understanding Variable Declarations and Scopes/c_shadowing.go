package main

import "fmt"

// c. Shadowing

// Go allows variables in inner scopes to have the same name as those in outer scopes, a concept known as shadowing.
// While powerful, it should be used judiciously to avoid confusion.

// Best Practice: Avoid unnecessary shadowing as it can make the code harder to read and debug.
// Use distinct variable names when possible.

var scope = "Global"

func main() {
	scope := "Local"
	fmt.Println(scope) // Outputs: Local
	{
		scope := "Block"
		fmt.Println(scope) // Outputs: Block
	}
	fmt.Println(scope) // Outputs: Local
}
