package main

func maxSubarrayLength(nums []int, k int) int {
	n := len(nums)
	answer := 0
	left := 0
	counter := make(map[int]int)
	for right := 0; right < n; right++ {
		num := nums[right]
		counter[num]++
		count := counter[num]
		for left < right && count > k {
			if num == nums[left] {
				count--
			}
			counter[nums[left]]--
			left++
		}
		answer = max(answer, right-left+1)
	}
	return answer
}
