package main

func lexicalOrder(n int) []int {
	result := make([]int, 0)
	for i := 1; i <= 9; i++ {
		lexicalOrderDfs(i, n, &result)
	}
	return result
}

func lexicalOrderDfs(curr, n int, arr *[]int) {
	if curr > n {
		return
	}

	*arr = append(*arr, curr)

	for i := 0; i <= 9; i++ {
		next := curr*10 + i
		if next > n {
			break
		}
		lexicalOrderDfs(next, n, arr)
	}
}
