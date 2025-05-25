package main

import "testing"

func TestLongPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  int
	}{
		{
			name:  "example 1",
			words: []string{"lc", "cl", "gg"},
			want:  6,
		},
		{
			name:  "example 2",
			words: []string{"ab", "ty", "yt", "lc", "cl", "ab"},
			want:  8,
		},
		{
			name:  "example 3",
			words: []string{"cc", "ll", "xx"},
			want:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.words); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
