package main

import "testing"

func TestMaxNumberOfKSumPairTest(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 2, 3, 4}, 5, 2},
		{[]int{3, 1, 3, 4, 3}, 6, 1},
		{[]int{3, 5, 1, 5}, 2, 0},
		{[]int{2, 2, 2, 3, 1, 1, 4, 1}, 4, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxOperations(tt.nums, tt.k); got != tt.want {
				t.Errorf("maxOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
