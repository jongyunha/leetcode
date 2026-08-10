package main

type State struct {
	i int
	m int
}

func stoneGameII(piles []int) int {
	n := len(piles)
	suffixSum := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffixSum[i] = suffixSum[i+1] + piles[i]
	}

	memo := make(map[State]int)
	var dp func(i int, m int) (take int)
	dp = func(i int, m int) (take int) {
		if i >= n {
			return
		}

		if i+2*m >= n {
			return suffixSum[i]
		}

		state := State{i: i, m: m}
		if _, ok := memo[state]; ok {
			return memo[state]
		}

		best := 0
		for x := 1; x <= 2*m; x++ {
			nextM := max(x, m)
			opponent := dp(i+x, nextM)
			current := suffixSum[i] - opponent
			best = max(current, best)
		}
		memo[state] = best
		return best

	}
	return dp(0, 1)
}
