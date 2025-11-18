package easy

func kLengthApart(nums []int, k int) bool {
	prevIndex := -1
	for i, num := range nums {
		if num == 1 {
			if prevIndex == -1 {
				prevIndex = i
			} else {
				if i-prevIndex-1 >= k {
					prevIndex = i
				} else {
					return false
				}
			}
		}
	}

	return true
}
