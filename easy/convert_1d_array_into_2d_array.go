package easy

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return nil
	}

	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = original[i*n+j]
		}
	}

	return result
}
