package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum3(t *testing.T) {
	testcase := []struct {
		k        int
		n        int
		expected [][]int
	}{
		//{
		//	k: 3,
		//	n: 7,
		//	expected: [][]int{
		//		{1, 2, 4},
		//	},
		//},
		{
			k: 9,
			n: 45,
			expected: [][]int{
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
		},
	}

	for _, tc := range testcase {
		got := combinationSum3(tc.k, tc.n)
		assert.Equal(t, tc.expected, got)
	}
}
