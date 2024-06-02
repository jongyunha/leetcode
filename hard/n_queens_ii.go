package hard

func totalNQueens(n int) int {
	chessboard := make([][]int, n)
	for i := range chessboard {
		chessboard[i] = make([]int, n)
	}

	var cases int = 0

	var queenDfs func(level int, chessboard *[][]int)
	queenDfs = func(level int, chessboard *[][]int) {
		if level == n {
			cases++
			return
		}

		for j := range (*chessboard)[level] {
			if isAttackAble(chessboard, level, j) {
				(*chessboard)[level][j] = 1
				queenDfs(level+1, chessboard)
				(*chessboard)[level][j] = 0
			}
		}
	}

	queenDfs(0, &chessboard)
	return cases
}

func isAttackAble(chessboard *[][]int, x, y int) bool {
	for i := range *chessboard {
		for j := range (*chessboard)[i] {
			if i == x && (*chessboard)[i][j] == 1 {
				return false
			}
			if j == y && (*chessboard)[i][j] == 1 {
				return false
			}
			if i-j == x-y && (*chessboard)[i][j] == 1 {
				return false
			}
			if i+j == x+y && (*chessboard)[i][j] == 1 {
				return false
			}
		}
	}

	return true
}
