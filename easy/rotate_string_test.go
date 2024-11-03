package easy

import "testing"

// Given two strings s and goal, return true if and only if s can become goal after some number of shifts on s.
//
// A shift on s consists of moving the leftmost character of s to the rightmost position.
//
// For example, if s = "abcde", then it will be "bcdea" after one shift.
//
// Example 1:
//
// Input: s = "abcde", goal = "cdeab"
// Output: true
// Example 2:
//
// Input: s = "abcde", goal = "abced"
// Output: false
func Test_rotate_string(t *testing.T) {
	type args struct {
		s    string
		goal string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				s:    "abcde",
				goal: "cdeab",
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				s:    "abcde",
				goal: "abced",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateString(tt.args.s, tt.args.goal); got != tt.want {
				t.Errorf("rotate_string() = %v, want %v", got, tt.want)
			}
		})
	}
}
