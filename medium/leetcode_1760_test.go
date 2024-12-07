package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumSize(t *testing.T) {
	tests := []struct {
		nums          []int
		maxOperations int
		want          int
	}{
		{
			nums:          []int{9},
			maxOperations: 2,
			want:          3,
		},
		// {
		// 	nums:          []int{2, 4, 8, 2},
		// 	maxOperations: 4,
		// 	want:          2,
		// },
		// {
		// 	nums:          []int{7, 17},
		// 	maxOperations: 2,
		// 	want:          7,
		// },
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, minimumSize(tt.nums, tt.maxOperations))
	}
}
