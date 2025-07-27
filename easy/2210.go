package easy

func countHillValley(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	n := len(nums)
	count := 0
	i := 1

	for i < n-1 {
		left := i - 1
		for left >= 0 && nums[left] == nums[i] {
			left--
		}

		right := i + 1
		for right < n && nums[right] == nums[i] {
			right++
		}

		if left >= 0 && right < n {
			leftVal := nums[left]
			rightVal := nums[right]
			currVal := nums[i]

			isValley := leftVal > currVal && rightVal > currVal
			isHill := leftVal < currVal && rightVal < currVal

			if isHill || isValley {
				count++
			}

			for i < n-1 && nums[i] == nums[i+1] {
				i++
			}
		}

		i++
	}

	return count
}
