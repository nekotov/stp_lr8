package main

import (
	"testing"
)

func TestMinPositive(t *testing.T) {
	alg := &Algorithms{}

	tests := []struct {
		input          []int
		expectedResult int
		expectError    bool
	}{
		{[]int{-3, 4, 2, -1}, 2, false},
		{[]int{-3, -4, -2, -1}, 0, true}, 		
		{[]int{5, 7, 1, 3}, 1, false},
	}

	for _, test := range tests {
		result, err := alg.MinPositive(test.input)
		if (err != nil) != test.expectError {
			t.Errorf("expected error: %v, got: %v", test.expectError, err)
		}
		if result != test.expectedResult {
			t.Errorf("expected: %d, got: %d", test.expectedResult, result)
		}
	}
}

func TestSumNegative(t *testing.T) {
	alg := &Algorithms{}

	tests := []struct {
		input          []int
		expectedResult int
	}{
		{[]int{-3, 4, -2, -1}, -6},
		{[]int{3, 4, 2, 1}, 0}, 		
		{[]int{-5, -7, -1, -3}, -16},
	}

	for _, test := range tests {
		result := alg.SumNegative(test.input)
		if result != test.expectedResult {
			t.Errorf("expected: %d, got: %d", test.expectedResult, result)
		}
	}
}

func TestFibonacci(t *testing.T) {
	alg := &Algorithms{}

	tests := []struct {
		input          int
		expectedResult int
		expectError    bool
	}{
		{10, 55, false},
		{-1, 0, true}, 		{0, 0, false},
		{1, 1, false},
	}

	for _, test := range tests {
		result, err := alg.Fibonacci(test.input)
		if (err != nil) != test.expectError {
			t.Errorf("expected error: %v, got: %v", test.expectError, err)
		}
		if result != test.expectedResult {
			t.Errorf("expected: %d, got: %d", test.expectedResult, result)
		}
	}
}

func TestCurrent(t *testing.T) {
	alg := &Algorithms{}

	tests := []struct {
		voltage        float64
		resistance     float64
		expectedResult float64
		expectError    bool
	}{
		{12, 4, 3, false},
		{5, 0, 0, true}, 		
		{10, 2, 5, false},
	}

	for _, test := range tests {
		result, err := alg.Current(test.voltage, test.resistance)
		if (err != nil) != test.expectError {
			t.Errorf("expected error: %v, got: %v", test.expectError, err)
		}
		if result != test.expectedResult {
			t.Errorf("expected: %f, got: %f", test.expectedResult, result)
		}
	}
}
