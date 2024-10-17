package main

import "fmt"

// Defining methods on custom types enhances functionality and encapsulation.

type Celsius float64

func (c Celsius) ToFahrenheit() float64 {
	return float64(c)*9/5 + 32
}

func main() {
	temp := Celsius(25)
	fmt.Printf("Temperature: %.2f°F\n", temp.ToFahrenheit())
}
