package main

import "fmt"

func arrays() {
	// numbers := [10]int{100, 123, 1, 15, 53, 213, 43, 23, 432, 12}
	// for _, number := range numbers {
	// 	fmt.Printf("%d \n", number)
	// }
	names := [5]string{"Alice", "Audery", "Klein", "Jack", "Gerhman"}
	for _, name := range names {
		fmt.Printf("Welcome to the venue, %s hope you have wonderful day! \n", name)
	}
	floatingNums := [3]float64{198.12, 2.35, 7.09}
	total := 0.0
	for _, number := range floatingNums {
		total += number
	}
	fmt.Printf("total = %f \n", total)
}
