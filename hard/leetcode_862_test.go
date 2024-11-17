package hard

import "testing"

// Given an integer array nums and an integer k, return the length of the shortest non-empty subarray of nums with a sum of at least k. If there is no such subarray, return -1.
//
// A subarray is a contiguous part of an array.
//
// Example 1:
//
// Input: nums = [1], k = 1
// Output: 1
// Example 2:
//
// Input: nums = [1,2], k = 4
// Output: -1
// Example 3:
//
// Input: nums = [2,-1,2], k = 3
// Output: 3
func TestShortestSubarray(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{nums: []int{1}, k: 1, want: 1},
		//{nums: []int{1, 2}, k: 4, want: -1},
		//{nums: []int{2, -1, 2}, k: 3, want: 3},
		//{nums: []int{17, 85, 93, -45, -21}, k: 150, want: 2},
		//{nums: []int{-28, 81, -20, 28, -29}, k: 89, want: 3},
	}

	for _, tt := range tests {
		if got := shortestSubarray(tt.nums, tt.k); got != tt.want {
			t.Errorf("ShortestSubArray(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
		}
	}
}
