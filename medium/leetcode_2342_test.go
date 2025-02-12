package main

import "testing"

// You are given a 0-indexed array nums consisting of positive integers. You can choose two indices i and j, such that i != j, and the sum of digits of the number nums[i] is equal to that of nums[j].
//
//Return the maximum value of nums[i] + nums[j] that you can obtain over all possible indices i and j that satisfy the conditions.
//
//
//
//Example 1:
//
//Input: nums = [18,43,36,13,7]
//Output: 54
//Explanation: The pairs (i, j) that satisfy the conditions are:
//- (0, 2), both numbers have a sum of digits equal to 9, and their sum is 18 + 36 = 54.
//- (1, 4), both numbers have a sum of digits equal to 7, and their sum is 43 + 7 = 50.
//So the maximum sum that we can obtain is 54.
//Example 2:
//
//Input: nums = [10,12,19,14]
//Output: -1
//Explanation: There are no two numbers that satisfy the conditions, so we return -1.

func TestMaximumSum(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		//{nums: []int{18, 43, 36, 13, 7}, expected: 54},
		//{nums: []int{10, 12, 19, 14}, expected: -1},
		{nums: []int{229, 398, 269, 317, 420, 464, 491, 218, 439, 153, 482, 169, 411, 93, 147, 50, 347, 210, 251, 366, 401}, expected: 973},
	}

	for _, tc := range cases {
		output := maximumSum(tc.nums)
		if output != tc.expected {
			t.Fatalf("expected: %v, got: %v", tc.expected, output)
		}
	}
}
