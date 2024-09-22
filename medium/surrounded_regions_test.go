package main

import (
	"fmt"
	"testing"
)

func Test_surrounded_regions(t *testing.T) {
	type args struct {
		board    [][]byte
		expected [][]byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test 1",
			args: args{
				board: [][]byte{
					{'X', 'X', 'X', 'X'},
					{'X', 'O', 'O', 'X'},
					{'X', 'X', 'O', 'X'},
					{'X', 'O', 'X', 'X'},
				},
				expected: [][]byte{
					{'X', 'X', 'X', 'X'},
					{'X', 'X', 'X', 'X'},
					{'X', 'X', 'X', 'X'},
					{'X', 'O', 'X', 'X'},
				},
			},
		},
		{
			name: "Test 2",
			args: args{
				board: [][]byte{
					{'X', 'X', 'X', 'X'},
					{'X', 'O', 'O', 'X'},
					{'X', 'X', 'O', 'X'},
					{'X', 'X', 'X', 'X'},
				},
				expected: [][]byte{
					{'X', 'X', 'X', 'X'},
					{'X', 'X', 'X', 'X'},
					{'X', 'X', 'X', 'X'},
					{'X', 'X', 'X', 'X'},
				},
			},
		},
		{
			name: "Test 3",
			args: args{
				board: [][]byte{
					{'X', 'O', 'X'},
					{'O', 'X', 'O'},
					{'X', 'O', 'X'},
				},
				expected: [][]byte{
					{'X', 'O', 'X'},
					{'O', 'X', 'O'},
					{'X', 'O', 'X'},
				},
			},
		},
		{
			name: "Test 4",
			args: args{
				board: [][]byte{
					{'O', 'O', 'O', 'O', 'X', 'X'},
					{'O', 'O', 'O', 'O', 'O', 'O'},
					{'O', 'X', 'O', 'X', 'O', 'O'},
					{'O', 'X', 'O', 'O', 'X', 'O'},
					{'O', 'X', 'O', 'X', 'O', 'O'},
					{'O', 'X', 'O', 'O', 'O', 'O'},
				},
				expected: [][]byte{
					{'O', 'O', 'O', 'O', 'X', 'X'},
					{'O', 'O', 'O', 'O', 'O', 'O'},
					{'O', 'X', 'O', 'X', 'O', 'O'},
					{'O', 'X', 'O', 'O', 'X', 'O'},
					{'O', 'X', 'O', 'X', 'O', 'O'},
					{'O', 'X', 'O', 'O', 'O', 'O'},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.args.board)
			if !compareBoards(tt.args.board, tt.args.expected) {
				showDiff(tt.args.board, tt.args.expected)
				t.Errorf("Expected %v but got %v", tt.args.expected, tt.args.board)
			}
		})
	}
}

func compareBoards(board1, board2 [][]byte) bool {
	if len(board1) != len(board2) {
		return false
	}
	for i := range board1 {
		if len(board1[i]) != len(board2[i]) {
			return false
		}
		for j := range board1[i] {
			if board1[i][j] != board2[i][j] {
				return false
			}
		}
	}
	return true
}

func showDiff(board1, board2 [][]byte) {
	fmt.Println("Got:")
	for i := range board1 {
		for j := range board1[i] {
			fmt.Printf("%c ", board1[i][j])
		}
		fmt.Println()
	}

	fmt.Println("Expected:")
	for i := range board2 {
		for j := range board2[i] {
			fmt.Printf("%c ", board2[i][j])
		}
		fmt.Println()
	}
}
