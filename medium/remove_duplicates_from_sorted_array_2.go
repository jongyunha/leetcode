package main

func removeDuplicates(nums []int) int {
	mostAppear := 2
	left, right := 0, len(nums)-1
	duplicate := 0

	for left < right {
		if nums[left] == nums[left+1] {
			duplicate++
		} else if nums[left] != nums[left+1] {
			duplicate = 0
		}

		if duplicate >= mostAppear {
			for i := left; i < right-1; i++ {
				swap(i+1, i+2, nums)
			}
			right--
			continue
		}

		left++
	}

	return right + 1
}

func swap(i, j int, nums []int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
