package main

import "testing"

// You are given an integer n. We reorder the digits in any order (including the original order) such that the leading digit is not zero.
//
// Return true if and only if we can do this so that the resulting number is a power of two.
//
// Example 1:
//
// Input: n = 1
// Output: true
// Example 2:
//
// Input: n = 10
// Output: false
func TestReorderedPowerOf2(t *testing.T) {
	tests := []struct {
		n      int
		expect bool
	}{
		{1, true},
		{2, true},
		{10, false},
		{16, true},
		{24, false},
		{46, true},
		{218, true},
	}

	for _, test := range tests {
		result := reorderedPowerOf2(test.n)
		if result != test.expect {
			t.Errorf("reorderedPowerOf2(%d) = %v; want %v", test.n, result, test.expect)
		}
	}
}
