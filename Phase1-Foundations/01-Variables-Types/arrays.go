package main

import (
	"fmt"
	"slices"
)

func arrays() {
	// 1. Array Iteration
	names := [5]string{"Alice", "Audery", "Klein", "Jack", "Gerhman"}
	for _, name := range names {
		fmt.Printf("Welcome to the venue, %s! Hope you have a wonderful day!\n", name)
	}

	// 2. Summing values
	floatingNums := [3]float64{198.12, 2.35, 7.09}
	total := 0.0
	for _, number := range floatingNums {
		total += number
	}
	fmt.Printf("Total = %.2f\n", total)

	// 3. Slice Manipulation (Insertion)
	sliceArr := []string{"hello", "world"}
	s := slices.Insert(sliceArr, 2, "morning")
	for _, word := range s {
		fmt.Printf("%s\n", word)
	}

	// 4. Sorting
	// Note: We use a slice here because sorting usually implies dynamic data
	numbers := []int{10, 2, 34, 5, 66, 7}
	slices.Sort(numbers)

	fmt.Println("Sorted numbers:")
	for _, number := range numbers {
		fmt.Printf("%d ", number)
	}
	fmt.Println()
}
