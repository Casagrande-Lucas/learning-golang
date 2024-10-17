package main

import "fmt"

// While basic usage is covered in introductory topics,
// advanced manipulation and understanding of underlying mechanics are essential.

func main() {
	// Arrays: Fixed-size sequences of elements of the same type.
	var arr [5]int
	arr[0] = 10
	fmt.Println(arr)

	// Slices: Dynamic, flexible views into arrays.
	// Advanced Operations:

	// Capacity and Length Management
	s := []int{1, 2, 3}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// Appending and Copying Slices
	s = append(s, 4, 5)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// Using Slices in Functions
	s2 := make([]int, 3)
	copy(s2, s)
	fmt.Println(s2)
}
