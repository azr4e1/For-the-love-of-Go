// Package calculator does simple calculations.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the result of adding
// them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers a and b, and
// returns the result of subtracting b from a.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers a and b, and
// returns the result of multiplying a and b.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers a and b, and
// returns the result of dividing a with b.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero no allowed")
	}
	return a / b, nil
}

// SquareRoot takes one number a, and
// returns the result of calculating its square root.
func SquareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("square root of negative numbers not allowed.")
	}
	return math.Sqrt(a), nil
}

func AddMany(inputs ...float64) float64 {
	var total float64
	for _, input := range inputs {
		total += input
	}

	return total
}
