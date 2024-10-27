package main

import "testing"

func Test_remove_duplicate_letters_test(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "test 1",
			s:    "bcabc",
			want: "abc",
		},
		{
			name: "test 2",
			s:    "cbacdcbc",
			want: "acdb",
		},
		{
			name: "test 3",
			s:    "abacb",
			want: "abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicateLetters(tt.s); got != tt.want {
				t.Errorf("removeDuplicateLetters() = %v, want %v", got, tt.want)
			}
		})
	}

}
