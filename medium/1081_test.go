package main

import "testing"

// Given a string s, return the lexicographically smallest subsequence of s that contains all the distinct characters of s exactly once.
//
// Example 1:
//
// Input: s = "bcabc"
// Output: "abc"
// Example 2:
//
// Input: s = "cbacdcbc"
// Output: "acdb"
//
// Constraints:
//
// 1 <= s.length <= 1000
// s consists of lowercase English letters.
func TestSmallestSubsequence(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "test case 1",
			s:    "bcabc",
			want: "abc",
		},
		{
			name: "test case 2",
			s:    "cbacdcbc",
			want: "acdb",
		},
		{
			name: "test case 3",
			s:    "cdadabcc",
			want: "adbc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestSubsequence(tt.s); got != tt.want {
				t.Errorf("smallestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
