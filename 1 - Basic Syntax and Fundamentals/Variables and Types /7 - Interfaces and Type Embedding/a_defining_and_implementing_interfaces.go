package main

import "fmt"

// While interfaces are typically covered in intermediate topics, their interaction with types is essential.

// Interfaces define a set of method signatures that types must implement.

type Stringer interface {
	String() string
}

type Person struct {
	Name string
}

func (p Person) String() string {
	return p.Name
}

func printString(s Stringer) {
	fmt.Println(s.String())
}

func main() {
	p := Person{Name: "Diana"}
	printString(p)
}
