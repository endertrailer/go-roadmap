package main

import "errors"

type Rectangle struct {
	Length float64
	Width  float64
}

var squareError = errors.New("length or width less than or equal to zero")

func squareRectangle(width float64, length float64) (float64, error) {
	if width <= 0 || length <= 0 {
		return 0, squareError
	}
	return width * length, nil
}
