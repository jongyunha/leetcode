package main

import (
	"reflect"
	"testing"
)

// Given a set of distinct positive integers nums, return the largest subset answer such that every pair (answer[i], answer[j]) of elements in this subset satisfies:
//
// answer[i] % answer[j] == 0, or
// answer[j] % answer[i] == 0
// If there are multiple solutions, return any of them.
//
// Example 1:
//
// Input: nums = [1,2,3]
// Output: [1,2]
// Explanation: [1,3] is also accepted.
// Example 2:
//
// Input: nums = [1,2,4,8]
// Output: [1,2,4,8]
func TestLargestDivisibleSubset(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		//{[]int{1, 2, 3}, []int{1, 2}},
		//{[]int{1, 2, 4, 8}, []int{1, 2, 4, 8}},
		//{[]int{3, 4, 16, 8}, []int{4, 8, 16}},
		{[]int{2, 3, 4, 8}, []int{2, 4, 8}},
	}

	for i, tt := range tests {
		got := largestDivisibleSubset(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d. got %v, want %v", i, got, tt.want)
		}
	}
}
