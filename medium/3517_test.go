package main

import "testing"

// You are given a palindromic string s.
//
// Return the lexicographically smallest palindromic permutation of s.
//
// Example 1:
//
// Input: s = "z"
//
// Output: "z"
//
// Explanation:
//
// A string of only one character is already the lexicographically smallest palindrome.
//
// Example 2:
//
// Input: s = "babab"
//
// Output: "abbba"
//
// Explanation:
//
// Rearranging "babab" → "abbba" gives the smallest lexicographic palindrome.
//
// Example 3:
//
// Input: s = "daccad"
//
// Output: "acddca"
//
// Explanation:
//
// Rearranging "daccad" → "acddca" gives the smallest lexicographic palindrome.
func TestSmallestPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"z", "z"},
		{"babab", "abbba"},
		{"daccad", "acddca"},
		{"rur", "rur"},
		{"eyy", "yey"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := smallestPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}
