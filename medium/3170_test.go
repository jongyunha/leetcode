package main

import "testing"

func TestClearStars(t *testing.T) {
	tests := []struct {
		s        string
		expected string
	}{
		{"aaba*", "aab"},
		{"abc", "abc"},
		{"d*d*", ""},
		{"dk**", ""},
		{"d*o*", ""},
	}

	for _, test := range tests {
		result := clearStars(test.s)
		if result != test.expected {
			t.Errorf("clearStars(%q) = %q; want %q", test.s, result, test.expected)
		}
	}
}
