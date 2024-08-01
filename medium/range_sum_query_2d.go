package main

type NumMatrix struct {
	matrix    [][]int
	sumMatrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	sumMatrix := make([][]int, len(matrix))

	for i := range sumMatrix {
		sumMatrix[i] = make([]int, len(matrix[i]))
	}

	sumMatrix[0][0] = matrix[0][0]

	for i := 0; i < len(sumMatrix); i++ {
		for j := 0; j < len(sumMatrix[i]); j++ {
			if i == 0 && j == 0 {
				continue
			}

			if j == 0 {
				sumMatrix[i][j] = sumMatrix[i-1][len(sumMatrix[i])-1] + matrix[i][j]
			} else {
				sumMatrix[i][j] = sumMatrix[i][j-1] + matrix[i][j]
			}
		}
	}

	return NumMatrix{
		matrix:    matrix,
		sumMatrix: sumMatrix,
	}
}

func (m *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := m.sumMatrix[row2][col2]

	if row1 > 0 {
		sum -= m.sumMatrix[row1-1][len(m.sumMatrix[row1-1])-1]
	}

	var sub int

	for i := row1; i <= row2; i++ {
		for j := 0; j < col1; j++ {
			sub += m.matrix[i][j]
		}
	}

	for i := row1; i <= row2-1; i++ {
		for j := col2 + 1; j < len(m.matrix[i]); j++ {
			sub += m.matrix[i][j]
		}
	}

	return sum - sub
}
