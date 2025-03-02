package easy

func applyOperations(nums []int) []int {
	result := make([]int, 0)
	zeros := make([]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] = nums[i] * 2
			nums[i+1] = 0
		}
	}

	for _, num := range nums {
		if num != 0 {
			result = append(result, num)
		} else {
			zeros = append(zeros, num)
		}
	}

	return append(result, zeros...)
}
