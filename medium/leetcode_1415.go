package main

func getHappyString(n int, k int) string {
	result := make([]string, 0)
	var generateHappyString func(carry string, n int)
	generateHappyString = func(carry string, n int) {
		if len(carry) == n {
			result = append(result, carry)
			return
		}

		for _, s := range []string{"a", "b", "c"} {
			if string(carry[len(carry)-1]) == s {
				continue
			}
			generateHappyString(carry+s, n)
		}
	}
	for _, s := range []string{"a", "b", "c"} {
		generateHappyString(s, n)
		if len(result) >= k {
			return result[k-1]
		}
	}

	return ""
}
