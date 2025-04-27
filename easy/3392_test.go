package easy

import "testing"

// Given an integer array nums, return the number of subarrays of length 3 such that the sum of the first and third numbers equals exactly half of the second number.
//
// Example 1:
//
// Input: nums = [1,2,1,4,1]
//
// Output: 1
//
// Explanation:
//
// Only the subarray [1,4,1] contains exactly 3 elements where the sum of the first and third numbers equals half the middle number.
//
// Example 2:
//
// Input: nums = [1,1,1]
//
// Output: 0
//
// Explanation:
//
// [1,1,1] is the only subarray of length 3. However, its first and third numbers do not add to half the middle number.
func TestCountSubArrays(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 1, 4, 1}, 1},
		{[]int{1, 1, 1}, 0},
		{[]int{-1, -4, -1, 4}, 1},
	}

	for _, tt := range tests {
		t.Run("countSubarrays", func(t *testing.T) {
			if got := countSubarrays(tt.nums); got != tt.want {
				t.Errorf("countSubarrays(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}
