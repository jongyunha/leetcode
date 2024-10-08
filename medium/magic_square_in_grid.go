package main

func numMagicSquaresInside(grid [][]int) int {
	if len(grid) < 3 || len(grid[0]) < 3 {
		return 0
	}
	var count int
	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid[0])-2; j++ {
			if isMagicSquare(grid, i, j) {
				count++
			}
		}
	}
	return count
}

func isMagicSquare(grid [][]int, i, j int) bool {
	if grid[i+1][j+1] != 5 {
		return false
	}
	count := make([]int, 16)
	for x := i; x < i+3; x++ {
		for y := j; y < j+3; y++ {
			count[grid[x][y]]++
		}
	}
	for k := 1; k <= 9; k++ {
		if count[k] != 1 {
			return false
		}
	}
	if grid[i][j]+grid[i][j+1]+grid[i][j+2] != 15 {
		return false
	}
	if grid[i+1][j]+grid[i+1][j+1]+grid[i+1][j+2] != 15 {
		return false
	}
	if grid[i+2][j]+grid[i+2][j+1]+grid[i+2][j+2] != 15 {
		return false
	}
	if grid[i][j]+grid[i+1][j]+grid[i+2][j] != 15 {
		return false
	}
	if grid[i][j+1]+grid[i+1][j+1]+grid[i+2][j+1] != 15 {
		return false
	}
	if grid[i][j+2]+grid[i+1][j+2]+grid[i+2][j+2] != 15 {
		return false
	}
	if grid[i][j]+grid[i+1][j+1]+grid[i+2][j+2] != 15 {
		return false
	}
	if grid[i+2][j]+grid[i+1][j+1]+grid[i][j+2] != 15 {
		return false
	}
	return true
}
