package core

import (
	"errors"
	"math"
	"testing"
)

func approxEqual(a, b, tol float64) bool {
	return math.Abs(a-b) <= tol
}

func TestAddition(t *testing.T) {
	tests := []struct {
		a, b   float64
		expect float64
	}{
		{2, 3, 5},
		{1.5, 2.35, 3.85},
		{-5, -3, -8},
		{0, 7, 7},
		{99999999, 1, 100000000},
	}

	for _, test := range tests {
		result := Add(test.a, test.b)
		if result != test.expect {
			t.Errorf("Add(%v, %v) = %v; want %v", test.a, test.b, result, test.expect)
		}
	}
}

func TestSubtraction(t *testing.T) {
	tests := []struct {
		a, b   float64
		expect float64
	}{
		{10, 4, 6},
		{5.75, 2.25, 3.50},
		{-10, 7, -17},
		{15, 0, 15},
	}

	for _, test := range tests {
		result := Sub(test.a, test.b)
		if result != test.expect {
			t.Errorf("Sub(%v, %v) = %v; want %v", test.a, test.b, result, test.expect)
		}
	}
}

func TestMultiplication(t *testing.T) {
	tests := []struct {
		a, b   float64
		expect float64
	}{
		{7, 8, 56},
		{3.2, 2.5, 8.0},
		{-4, -6, 24},
		{0, 8, 0},
	}

	for _, test := range tests {
		result := Mul(test.a, test.b)
		if result != test.expect {
			t.Errorf("Mul(%v, %v) = %v; want %v", test.a, test.b, result, test.expect)
		}
	}
}

func TestDivision(t *testing.T) {
	type divCase struct {
		a, b      float64
		expect    float64
		expectErr bool
		approxTol float64
	}

	tests := []divCase{
		{20, 4, 5, false, 0},
		{10.5, 2.1, 5, false, 0},
		{-8, 2, -4, false, 0},
		// decimal precision checks
		{1, 3, 0.3333333333, false, 1e-9},
		{10, 6, 1.6666666667, false, 1e-9},
		// division by zero case
		{5, 0, 0, true, 0},
	}

	for _, tc := range tests {
		result, err := Div(tc.a, tc.b)

		if tc.expectErr {
			if err == nil {
				t.Errorf("Div(%v, %v) expected error but got none", tc.a, tc.b)
			}
			continue
		}

		if err != nil {
			t.Errorf("Div(%v, %v) unexpected error: %v", tc.a, tc.b, err)
			continue
		}

		if tc.approxTol > 0 {
			if !approxEqual(result, tc.expect, tc.approxTol) {
				t.Errorf("Div(%v, %v) = %v; want approx %v", tc.a, tc.b, result, tc.expect)
			}
		} else if result != tc.expect {
			t.Errorf("Div(%v, %v) = %v; want %v", tc.a, tc.b, result, tc.expect)
		}
	}
}
