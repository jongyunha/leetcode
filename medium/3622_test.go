package main

import "testing"

func TestCheckDivisibility(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{
			n:    8,
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run("test", func(t *testing.T) {
			if checkDivisibility(tc.n) != tc.want {
				t.Fail()
			}
		})
	}
}
