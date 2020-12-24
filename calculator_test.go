package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name string
		a, b float64
		want float64
	}

	testCases := []testCase{
		{name: "Standard adding", a: 10, b: 4, want: 14},
		{name: "Negative added onto positive", a: 4, b: -8, want: -4},
		{name: "Adding decimals", a: 15.5, b: 4.5, want: 20},
		{name: "Positive added onto negative", a: -5, b: 10, want: 5},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)

		if tc.want != got {
			t.Errorf("Adding (%f and %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestAddRandom(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{}

	for i := 0; i < 100; i++ {
		a, b := rand.Float64(), rand.Float64()
		var want float64 = a + b
		testCases = append(testCases, testCase{a: a, b: b, want: want})
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Adding (%f and %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}

}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name string
		a, b float64
		want float64
	}

	testCases := []testCase{
		{name: "Standard subtraction", a: 10, b: 4, want: 6},
		{name: "Going below zero", a: 4, b: 8, want: -4},
		{name: "Subtracting from decimals", a: 15.5, b: 4, want: 11.5},
		{name: "Subtracting two negatives to make a postive", a: -5, b: -10, want: 5},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)

		if tc.want != got {
			t.Errorf("%s (Subtracting (%f and %f)): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name string
		a, b float64
		want float64
	}

	testCases := []testCase{
		{name: "Standard multiplication", a: 10, b: 4, want: 40},
		{name: "Negativizing a number", a: 4, b: -8, want: -32},
		{name: "double negative multiplication", a: -15, b: -4, want: 60},
		{name: "Multiplying by 0", a: -5, b: 0, want: 0},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)

		if tc.want != got {
			t.Errorf("%s (Multiplying (%f and %f)): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name        string
		a, b        float64
		want        float64
		errExpected bool
	}

	testCases := []testCase{
		{name: "Standard Division", a: 8, b: 2, want: 4, errExpected: false},
		{name: "Dividing by negative", a: 16, b: -8, want: -2, errExpected: false},
		{name: "Dividing by zero", a: 5, b: 0, want: 0, errExpected: true},
		{name: "Dividing zero by number", a: 0, b: 5, want: 0, errExpected: false},
		{name: "Dividing decimals", a: 5.2, b: 0.5, want: 10.4, errExpected: false},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Fatalf("%s (%f and %f): Unexpected error status: %v", tc.name, tc.a, tc.b, errReceived)
		}
		if !tc.errExpected && tc.want != got {
			t.Fatalf("%s (%f and %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name        string
		a           float64
		want        float64
		errExpected bool
	}

	testCases := []testCase{
		{name: "Standard Sqrt", a: 64, want: 8, errExpected: false},
		{name: "Negative Sqrt", a: -16, want: 424242, errExpected: true},
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Fatalf("%s (%f): Unexpected error status: %v", tc.name, tc.a, errReceived)
		}
		if !tc.errExpected && tc.want != got {
			t.Fatalf("%s (%f): want %f, got %f", tc.name, tc.a, tc.want, got)
		}
	}
}
