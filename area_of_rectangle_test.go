package main

import (
	"testing"
)

func TestAreaOfRectangle(t *testing.T) {
	testCases := []struct {
		length   int
		breadth  int
		expected int
	}{
		{5, 6, 30},    // Test case with positive values
		{0, 0, 0},     // Test case with zero values
		{3, 4, 12},    // Test case with small values
		{10, 10, 100}, // Test case with equal values
		{-2, 3, -6},   // Test case with negative values
	}

	for _, tc := range testCases {
		result := AreaOfRectangle(tc.length, tc.breadth)
		if result != tc.expected {
			t.Errorf("AreaOfRectangle(%d, %d) = %d; expected %d", tc.length, tc.breadth, result, tc.expected)
		}
	}
}
