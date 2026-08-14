package easy

func maximumLengthSubstring(s string) int {
	left := 0
	n := len(s)
	counter := make(map[byte]int)
	result := 0
	for right := 0; right < n; right++ {
		c := s[right]
		counter[c]++
		if counter[c] <= 2 {
			result = max(result, right-left+1)
			continue
		}

		for counter[c] > 2 && left <= right {
			lc := s[left]
			counter[lc]--
			left++
		}
	}

	return result
}
