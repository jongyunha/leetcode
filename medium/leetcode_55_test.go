package main

import "testing"

// You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.
//
// Return true if you can reach the last index, or false otherwise.
//
// Example 1:
//
// Input: nums = [2,3,1,1,4]
// Output: true
// Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
// Example 2:
//
// Input: nums = [3,2,1,0,4]
// Output: false
// Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
func TestCanJump(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{0, 2, 3}, false},
		{[]int{0}, true},
	}
	for _, tt := range tests {
		if got := canJump(tt.nums); got != tt.want {
			t.Errorf("canJump(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
