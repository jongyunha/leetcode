package easy

import "testing"

// Given an binary array nums and an integer k, return true if all 1's are at least k places away from each other, otherwise return false.
// Example 1:
// Input: nums = [1,0,0,0,1,0,0,1], k = 2
// Output: true
// Explanation: Each of the 1s are at least 2 places away from each other.
// Example 2:
// Input: nums = [1,0,0,1,0,1], k = 2
// Output: false
// Explanation: The second 1 and third 1 are only one apart from each other.
func TestKLengthApart(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want bool
	}{
		{[]int{1, 0, 0, 0, 1, 0, 0, 1}, 2, true},
		{[]int{1, 0, 0, 1, 0, 1}, 2, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := kLengthApart(tt.nums, tt.k); got != tt.want {
				t.Errorf("kLengthApart(%v, %d) = %v; want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
