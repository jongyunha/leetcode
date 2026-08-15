package main

import (
	"fmt"
	"testing"
)

func TestLongestSubsequence(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{
			nums:   []int{1, 2, 3},
			output: 2,
		},
		{
			nums:   []int{2, 3, 4},
			output: 3,
		},
		{
			nums:   []int{7, 6, 1, 9},
			output: 4,
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			result := longestSubsequence(tc.nums)
			if result != tc.output {
				t.Errorf("want: %d, result: %d", tc.output, result)
			}
		})
	}
}
