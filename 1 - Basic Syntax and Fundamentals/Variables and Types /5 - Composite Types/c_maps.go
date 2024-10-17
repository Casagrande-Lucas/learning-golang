package main

import "fmt"

// Maps associate keys with values, offering efficient lookups.
// Best Practice: Use simple and comparable types as map keys to ensure efficient and predictable behavior.

func main() {
	//Advanced Usage:
	// • Concurrent Access with Sync Maps
	// • Custom Key Types
	// • Handling nil Maps

	m := make(map[string]int)
	m["apple"] = 5
	m["banana"] = 3

	for key, value := range m {
		fmt.Printf("%s: %d\n", key, value)
	}

	// Using struct as a key
	type Coordinate struct {
		X, Y int
	}

	coordMap := map[Coordinate]string{
		Coordinate{0, 0}: "Origin",
		Coordinate{1, 2}: "Point A",
	}

	fmt.Println(coordMap)
}
