package main

import "testing"

func Test_magic_square_in_grid_test(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "test 1",
			grid: [][]int{
				{4, 3, 8, 4},
				{9, 5, 1, 9},
				{2, 7, 6, 2},
			},
			want: 1,
		},
		{
			name: "test 2",
			grid: [][]int{
				{8},
			},
			want: 0,
		},
		{
			name: "edge case 1",
			grid: [][]int{
				{5, 5, 5},
				{5, 5, 5},
				{5, 5, 5},
			},
			want: 0,
		},
		{
			name: "edge case 2",
			grid: [][]int{
				{3, 2, 9, 2, 7},
				{6, 1, 8, 4, 2},
				{7, 5, 3, 2, 7},
				{2, 9, 4, 9, 6},
				{4, 3, 8, 2, 5},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numMagicSquaresInside(tt.grid); got != tt.want {
				t.Errorf("magic_square_in_grid() = %v, want %v", got, tt.want)
			}
		})
	}

}
