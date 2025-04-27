package easy

func countSubarrays(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	count := 0
	for i := 0; i < len(nums)-2; i++ {
		first, second, third := nums[i], nums[i+1], nums[i+2]

		sum := float64(first + third)
		half := float64(second) / float64(2)

		if sum == half {
			count++
		}
	}

	return count
}
