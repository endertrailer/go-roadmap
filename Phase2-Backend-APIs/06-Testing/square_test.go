package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	type testCase struct {
		name     string
		height   float64
		width    float64
		expected float64
		err      error
	}

	tests := []testCase{
		{name: "Standard Integer Bounds", height: 10, width: 7, expected: 70},
		{name: "Asymmetric Bounds", height: 20, width: 2, expected: 40},
		{name: "Floating Point Precision", height: 8.5, width: 6.7, expected: 56.95},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := squareRectangle(test.height, test.width)

			assert.NoError(t, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}
