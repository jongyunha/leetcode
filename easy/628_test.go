package easy

import "testing"

// Given an integer array nums, find three numbers whose product is maximum and return the maximum product.
//
// Example 1:
//
// Input: nums = [1,2,3]
// Output: 6
// Example 2:
//
// Input: nums = [1,2,3,4]
// Output: 24
// Example 3:
//
// Input: nums = [-1,-2,-3]
// Output: -6
//
// Constraints:
//
// 3 <= nums.length <= 104
// -1000 <= nums[i] <= 1000
func TestMaximumProduct(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{1, 2, 3, 4}, 24},
		{[]int{-1, -2, -3}, -6},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := maximumProduct(tt.nums)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}
