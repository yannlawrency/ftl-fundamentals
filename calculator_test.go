package calculator_test

import (
	"calculator"
	"testing"
)

type testCase struct {
	a, b float64
	want float64
	name string
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{name: "add two integers", a: 2, b: 2, want: 4},
		{name: "add floats", a: 0.1, b: 0.9, want: 1},
		{name: "add a negative numbers", a: -10, b: 9, want: -1},
		{name: "add two negative numbers", a: -10, b: -9, want: -19},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s: (%f+%f) -> want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{name: "subtract two integers", a: 2, b: 2, want: 0},
		{name: "subtract floats", a: 1.2, b: 0.3, want: 0.9},
		{name: "subtract a negative number", a: -10, b: 9, want: -19},
		{name: "subtract two negative numbers", a: -10, b: -9, want: -1},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s: (%f-%f) -> want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{name: "multiply two integers", a: 2, b: 3, want: 6},
		{name: "multiply a negative number", a: -10, b: 9, want: -90},
		{name: "multiply two negative numbers", a: -10, b: -9, want: 90},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s: (%f*%f) -> want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}
