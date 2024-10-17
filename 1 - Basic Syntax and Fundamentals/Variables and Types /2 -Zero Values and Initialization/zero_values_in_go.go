package main

import "fmt"

// Zero Values in Go

// • Numeric Types: 0
// • Boolean: false
// • String: "" (empty string)
// • Pointers, Functions, Interfaces, Slices, Channels, Maps: nil

// Best Practice: Explicitly initialize variables when zero values are not appropriate or to improve code readability.

var (
	num   int
	flag  bool
	text  string
	ptr   *int
	slice []int
	myMap map[string]int
)

func main() {
	fmt.Printf("num: %d, flag: %t, text: '%s'\n", num, flag, text)
	fmt.Printf("ptr: %v, slice: %v, myMap: %v\n", ptr, slice, myMap)
}
