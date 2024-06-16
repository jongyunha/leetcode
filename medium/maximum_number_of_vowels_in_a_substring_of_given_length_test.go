package main

import "testing"

func Test_maximum_number_of_vowels_in_a_substring_of_given_length(t *testing.T) {
	type args struct {
		s      string
		k      int
		output int
	}
	tests := []args{
		{"abciiidef", 3, 3},
		{"aeiou", 2, 2},
		{"leetcode", 3, 2},
		{"rhythms", 4, 0},
		{"tryhard", 4, 1},
		{"weallloveyou", 7, 4},
		{"ibpbhixfiouhdljnjfflpapptrxgcomvnb", 33, 7},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := maxVowels(tt.s, tt.k); got != tt.output {
				t.Errorf("maximum_number_of_vowels_in_a_substring_of_given_length() = %v, want %v", got, tt.output)
			}
		})
	}
}
