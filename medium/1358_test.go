package main

import "testing"

// Given a string s consisting only of characters a, b and c.
//
// Return the number of substrings containing at least one occurrence of all these characters a, b and c.
//
// Example 1:
//
// Input: s = "abcabc"
// Output: 10
// Explanation: The substrings containing at least one occurrence of the characters a, b and c are "abc", "abca", "abcab", "abcabc", "bca", "bcab", "bcabc", "cab", "cabc" and "abc" (again).
// Example 2:
//
// Input: s = "aaacb"
// Output: 3
// Explanation: The substrings containing at least one occurrence of the characters a, b and c are "aaacb", "aacb" and "acb".
// Example 3:
//
// Input: s = "abc"
// Output: 1
func TestNumberOfSubstrings(t *testing.T) {
	cases := []struct {
		s      string
		expect int
	}{
		{"abcabc", 10},
		//{"aaacb", 3},
		//{"abc", 1},
	}

	for _, data := range cases {
		if result := numberOfSubstrings(data.s); result != data.expect {
			t.Errorf("expect %d, but got %d", data.expect, result)
		}
	}
}
