package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContinuousSubarrys(t *testing.T) {
	tests := []struct {
		nums []int
		want int64
	}{
		// {
		// 	nums: []int{5, 4, 2, 4},
		// 	want: 8,
		// },
		// {
		// 	nums: []int{1, 2, 3},
		// 	want: 6,
		// },
		// {
		// 	nums: []int{31, 30, 31, 32},
		// 	want: 10,
		// },
		{
			nums: []int{65, 66, 67, 66, 66, 65, 64, 65, 65, 64},
			want: 43,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, continuousSubarrays(tt.nums), tt.want)
	}
}
