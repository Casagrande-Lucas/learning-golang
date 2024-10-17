package main

import "fmt"

// Go offers several composite types that allow grouping of values.
// Structs are custom data types that group related fields.

func main() {
	type Person struct {
		Name   string
		Age    int
		Emails []string
	}

	p := Person{
		Name:   "Bob",
		Age:    25,
		Emails: []string{"bob@example.com", "bob.work@example.com"},
	}
	fmt.Println(p)

	// Advanced Usage:
	// Anonymous Fields: Embedding structs within structs.
	// Struct Tags: Adding metadata for serialization or validation.

	type Address struct {
		Street string `json:"street"`
		City   string `json:"city"`
	}

	type Employee struct {
		Person
		Address
		Position string
	}

	e := Employee{
		Person: Person{
			Name:   "Charlie",
			Age:    28,
			Emails: []string{"charlie@example.com"},
		},
		Address: Address{
			Street: "123 Go Lane",
			City:   "Gotham",
		},
		Position: "Developer",
	}
	fmt.Printf("%+v\n", e)
}
