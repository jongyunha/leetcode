package main

import "testing"

// You are given an array nums of non-negative integers and an integer k.
//
// An array is called special if the bitwise OR of all of its elements is at least k.
//
// Return the length of the shortest special non-empty
// subarray
// of nums, or return -1 if no special subarray exists.
//
// Example 1:
//
// Input: nums = [1,2,3], k = 2
//
// Output: 1
//
// Explanation:
//
// The subarray [3] has OR value of 3. Hence, we return 1.
//
// Example 2:
//
// Input: nums = [2,1,8], k = 10
//
// Output: 3
//
// Explanation:
//
// The subarray [2,1,8] has OR value of 11. Hence, we return 3.
//
// Example 3:
//
// Input: nums = [1,2], k = 0
//
// Output: 1
//
// Explanation:
//
// The subarray [1] has OR value of 1. Hence, we return 1.
func Test_shortest_subarray_with_or_at_least_k_ii(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		//{nums: []int{1, 2, 3}, k: 2, want: 1},
		//{nums: []int{2, 1, 8}, k: 10, want: 3},
		//{nums: []int{1, 2}, k: 0, want: 1},
		//{nums: []int{2, 89, 26, 12, 1, 62}, k: 99, want: 5},
		//{nums: []int{1, 2, 32, 21}, k: 55, want: 3},
		{nums: []int{2, 32, 27, 1}, k: 59, want: 2},
	}

	for _, tt := range tests {
		if got := minimumSubarrayLength(tt.nums, tt.k); got != tt.want {
			t.Errorf("sortest_subarray_with_or_at_least_k_ii(%v, %v) = %v, want %v", tt.nums, tt.k, got, tt.want)
		}
	}
}
