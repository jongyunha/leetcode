package main

import "testing"

func TestRemovingStarsFromString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		//{"leet**cod*e", "lecoe"},
		{"erase*****", ""},
	}

	for _, test := range tests {
		actual := removeStars(test.input)
		if actual != test.expected {
			t.Errorf("RemovingStarsFromString(%s): expected %s, actual %s", test.input, test.expected, actual)
		}
	}
}
