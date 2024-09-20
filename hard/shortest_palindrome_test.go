package hard

import "testing"

func Test_shortest_palindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Case 1", args{"aacecaaa"}, "aaacecaaa"},
		{"Case 2", args{"abcd"}, "dcbabcd"},
		{"Case 3", args{"a"}, "a"},
		{"Case 4", args{"aa"}, "aa"},
		{"Case 5", args{"abb"}, "bbabb"},
		{"Case 6", args{"abcd"}, "dcbabcd"},
		{"Case 7", args{"abbacd"}, "dcabbacd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("shortest_palindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
