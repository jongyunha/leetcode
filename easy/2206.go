package easy

func divideArray(nums []int) bool {
	counter := make(map[int]int)

	for _, num := range nums {
		counter[num]++
	}

	pair := counter[nums[0]]
	isPairEven := pair%2 == 0

	for _, value := range counter {
		if isPairEven && value%2 != 0 {
			return false
		} else if !isPairEven && value%3 != 0 {
			return false
		}
	}

	return true
}
