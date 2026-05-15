package main

import (
	"errors"
	"fmt"
)

func checkNumber(number int) (string, error) {
	if number < 0 {
		return "", errors.New("negative number")
	}
	return "Number is positive", nil
}

func main() {
	response, err := checkNumber(-199)
	if err == nil {
		fmt.Printf("%s \n", response)
	} else {
		println("numer is negative")
	}
}
