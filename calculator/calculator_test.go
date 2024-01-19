package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestAdd(t *testing.T) {
	t.Parallel()
	// define a struct for test cases
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 0},
		{a: 1, b: 4, want: -3},
		{a: -2, b: 0, want: -2},
		{a: 0, b: 0, want: 0},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 1},
		{a: 5, b: 0, want: 0},
		{a: -2, b: 4, want: -8},
		{a: -3, b: -5, want: 15},
		{a: 1, b: 1000, want: 1000},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	// define a struct for test cases
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 1},
		{a: -1, b: -1, want: 1},
		{a: 10, b: 2, want: 5},
		{a: 1, b: 3, want: 0.3333333333},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}
		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Divide(1, 0)
	if err == nil {
		t.Error("Want error for invalid input, got nil")
	}
}

func TestSquareRoot(t *testing.T) {
	t.Parallel()

	// define a struct for test cases
	type testCase struct {
		a    float64
		want float64
	}
	testCases := []testCase{
		{a: 4, want: 2},
		{a: 25, want: 5},
		{a: 16, want: 4},
		{a: 3, want: 1.73205081},
		{a: 0, want: 0},
	}

	for _, tc := range testCases {
		got, err := calculator.SquareRoot(tc.a)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}
		if !closeEnough(tc.want, got, 0.00001) {
			t.Errorf("SquareRoot(%f): want %f, got %f", tc.a, tc.want, got)
		}
	}
}

func TestSquareRootInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.SquareRoot(-1)
	if err == nil {
		t.Error("Want error for invalid input, got nil")
	}
}
