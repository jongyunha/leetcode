package easy

func divisorGame(n int) bool {
	memo := make(map[int]bool)
	var dp func(n int) bool

	dp = func(n int) bool {
		if n <= 1 {
			return false
		}

		if result, ok := memo[n]; ok {
			return result
		}

		for x := 1; x < n; x++ {
			if n%x != 0 {
				continue
			}

			if !dp(n - x) {
				memo[n] = true
				return true
			}
		}
		memo[n] = false
		return false
	}
	return dp(n)
}
