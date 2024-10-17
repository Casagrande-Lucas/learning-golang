package main

import "fmt"

// b. Custom Types and Underlying Types

// Creating custom types enhances type safety and expressiveness.

// Best Practice: Use custom types to represent distinct concepts,
// improving code clarity and preventing accidental misuse of values.

func main() {
	type Age int
	type Name string

	var myAge Age = 30
	var myName Name = "Alice"

	// var anotherAge int = myAge // Error: cannot use myAge (type Age) as type int
	var anotherAge int = int(myAge) // Correct

	// var anotherName string = myName // Error: cannot use myName (type Name) as the type string
	var anotherName string = string(myName) // Correct

	fmt.Println(anotherAge, anotherName)
}
