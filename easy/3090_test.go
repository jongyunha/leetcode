package easy

import (
	"testing"
)

func TestMaximumLengthSubString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "tc1",
			s:    "bcbbbcba",
			want: 4,
		},
		{
			name: "tc2",
			s:    "aaaa",
			want: 2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := maximumLengthSubstring(tc.s)
			if result != tc.want {
				t.Errorf("want: %d, actual: %d", tc.want, result)
			}
		})
	}
}
