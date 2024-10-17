package main

import "fmt"

// b. Variable Scope

// Variables in Go have specific scopes:
// • Package-Level Variables: Accessible throughout the package.
// • Function-Level Variables: Accessible only within the function.
// • Block-Level Variables: Accessible within the enclosing block (e.g. loops, conditionals).

var packageVar = "I am a package-level variable"

func main() {
	fmt.Println(packageVar)

	functionVar := "I am a function-level variable"
	fmt.Println(functionVar)

	if true {
		blockVar := "I am a block-level variable"
		fmt.Println(blockVar)
	}

	// fmt.Println(blockVar) // Error: undefined: blockVar
}
