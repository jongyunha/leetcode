package main

func lexicalOrder(n int) []int {
	res := make([]int, 0, n)
	for i := 1; i < 10; i++ {
		lexicalOrderDfs(i, n, &res)
	}
	return res
}

func lexicalOrderDfs(cur, n int, res *[]int) {
	if cur > n {
		return
	}
	*res = append(*res, cur)
	for i := 0; i < 10; i++ {
		lexicalOrderDfs(cur*10+i, n, res)
	}
}
