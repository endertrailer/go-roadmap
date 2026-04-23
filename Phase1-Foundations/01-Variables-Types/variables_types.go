package main

import (
	"fmt"
)

func variablestypes() {
	// how to declare a variable type is automatically inferred
	num := 10
	fmt.Printf("%d \n", num)
	// common types of variables
	number := 12345 // int

	text := "Phoenix A* is the biggest blackhole and has the mass roughly 100 billion times more than our sun" // string
	decimals := 3.141                                                                                          // float64
	print(number, "\n")
	print(text, "\n")
	print(decimals, "\n")
}
