package main

import "testing"

func Test_nearest_exit_from_entrance_in_maze(t *testing.T) {
	tests := []struct {
		maze     [][]byte
		entrance []int
		want     int
	}{
		{
			[][]byte{
				{'+', '+', '.', '+'},
				{'.', '.', '.', '+'},
				{'+', '+', '+', '.'},
			},
			[]int{1, 2},
			1,
		},
		//{
		//	[][]byte{
		//		{'+', '+', '+'},
		//		{'.', '.', '.'},
		//		{'+', '+', '+'},
		//	},
		//	[]int{1, 0},
		//	2,
		//},
		//{
		//	[][]byte{
		//		{'.', '+'},
		//	},
		//	[]int{0, 0},
		//	-1,
		//},
		//{
		//	[][]byte{
		//		{'.', '+', '.'},
		//	},
		//	[]int{0, 2},
		//	-1,
		//},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := nearestExit(tt.maze, tt.entrance); got != tt.want {
				t.Errorf("nearestExit() = %v, want %v", got, tt.want)
			}
		})
	}
}
