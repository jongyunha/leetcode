package main

// Given an integer array arr, remove a subarray (can be empty) from arr such that the remaining elements in arr are non-decreasing.
//
// Return the length of the shortest subarray to remove.
//
// A subarray is a contiguous subsequence of the array.
//
// Example 1:
//
// Input: arr = [1,2,3,10,4,2,3,5]
// Output: 3
// Explanation: The shortest subarray we can remove is [10,4,2] of length 3. The remaining elements after that will be [1,2,3,3,5] which are sorted.
// Another correct solution is to remove the subarray [3,10,4].
// Example 2:
//
// Input: arr = [5,4,3,2,1]
// Output: 4
// Explanation: Since the array is strictly decreasing, we can only keep a single element. Therefore we need to remove a subarray of length 4, either [5,4,3,2] or [4,3,2,1].
// Example 3:
//
// Input: arr = [1,2,3]
// Output: 0
// Explanation: The array is already non-decreasing. We do not need to remove any elements.
func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	left, right := 0, n-1

	for left < n-1 && arr[left] <= arr[left+1] {
		left++
	}

	if left == n-1 {
		return 0
	}

	for right > 0 && arr[right] >= arr[right-1] {
		right--
	}

	result := min(n-left-1, right)
	l, r := 0, right
	for l <= left && r < n {
		if arr[l] <= arr[r] {
			result = min(result, r-l-1)
			l++
		} else {
			r++
		}
	}

	return result
}
