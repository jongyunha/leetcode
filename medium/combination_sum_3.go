package main

func combinationSum3(k, n int) [][]int {
	carry := make([]int, 0)
	final := make([][]int, 0)

	backtracking(1, k, n, 0, carry, &final)

	return final
}

func backtracking(level, k, n, sum int, carry []int, final *[][]int) {
	if level > 10 {
		return
	}
	if sum == n && len(carry) == k {
		dst := make([]int, 0, len(carry))
		dst = append(dst, carry...)
		*final = append(*final, dst)
		return
	}

	if len(carry) >= k {
		return
	}

	for i := level; i < 10; i++ {
		carry = append(carry, i)
		sum += i
		backtracking(i+1, k, n, sum, carry, final)
		carry = carry[:len(carry)-1]
		sum -= i
	}
}
