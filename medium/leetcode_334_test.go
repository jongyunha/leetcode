package main

import "testing"

// Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.
//
// Example 1:
//
// Input: nums = [1,2,3,4,5]
// Output: true
// Explanation: Any triplet where i < j < k is valid.
// Example 2:
//
// Input: nums = [5,4,3,2,1]
// Output: false
// Explanation: No triplet exists.
// Example 3:
//
// Input: nums = [2,1,5,0,4,6]
// Output: true
// Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.
func TestIncreasingTriplet(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{nums: []int{1, 2, 3, 4, 5}, want: true},
		{nums: []int{5, 4, 3, 2, 1}, want: false},
		{nums: []int{2, 1, 5, 0, 4, 6}, want: true},
		{nums: []int{20, 100, 10, 12, 5, 13}, want: true},
		{nums: []int{4, 5, 2147483647, 1, 2}, want: true},
	}
	for _, tt := range tests {
		if got := increasingTriplet(tt.nums); got != tt.want {
			t.Errorf("increasingTriplet() = %v, want %v", got, tt.want)
		}
	}
}
