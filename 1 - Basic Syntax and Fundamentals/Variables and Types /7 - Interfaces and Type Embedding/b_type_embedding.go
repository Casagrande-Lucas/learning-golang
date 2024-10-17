package main

import "fmt"

// Embedding allows types to inherit fields and methods from other types, promoting code reuse.

// Best Practice: Use embedding to compose types and promote code reuse.
// Implement interfaces to define clear contracts for behavior.

type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Printf("%s makes a sound.\n", a.Name)
}

type Dog struct {
	Animal
	Breed string
}

func (d Dog) Speak() {
	fmt.Printf("%s barks.\n", d.Name)
}

func main() {
	d := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Golden Retriever",
	}
	d.Speak()        // Outputs: Buddy barks.
	d.Animal.Speak() // Outputs: Buddy makes a sound.
}
