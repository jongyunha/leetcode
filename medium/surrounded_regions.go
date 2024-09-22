package main

const (
	X = byte('X')
	O = byte('O')
	M = byte('M')
)

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		if board[i][0] == O {
			markBoard(i, 0, board)
		}
		if board[i][n-1] == O {
			markBoard(i, n-1, board)
		}
	}

	for j := 0; j < n; j++ {
		if board[0][j] == O {
			markBoard(0, j, board)
		}
		if board[m-1][j] == O {
			markBoard(m-1, j, board)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == O {
				board[i][j] = X
			}
			if board[i][j] == M {
				board[i][j] = O
			}
		}
	}
}

func markBoard(i, j int, board [][]byte) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != O {
		return
	}

	board[i][j] = M

	markBoard(i+1, j, board)
	markBoard(i-1, j, board)
	markBoard(i, j+1, board)
	markBoard(i, j-1, board)
}
