package main

import "testing"

// You are given two integers m and n representing a 0-indexed m x n grid. You are also given two 2D integer arrays guards and walls where guards[i] = [rowi, coli] and walls[j] = [rowj, colj] represent the positions of the ith guard and jth wall respectively.
//
// A guard can see every cell in the four cardinal directions (north, east, south, or west) starting from their position unless obstructed by a wall or another guard. A cell is guarded if there is at least one guard that can see it.
//
// Return the number of unoccupied cells that are not guarded.
//
// Example 1:
//
// Input: m = 4, n = 6, guards = [[0,0],[1,1],[2,3]], walls = [[0,1],[2,2],[1,4]]
// Output: 7
// Explanation: The guarded and unguarded cells are shown in red and green respectively in the above diagram.
// There are a total of 7 unguarded cells, so we return 7.
// Example 2:
//
// Input: m = 3, n = 3, guards = [[1,1]], walls = [[0,1],[1,0],[2,1],[1,2]]
// Output: 4
// Explanation: The unguarded cells are shown in green in the above diagram.
// There are a total of 4 unguarded cells, so we return 4.
func TestCountUnguarded(t *testing.T) {
	type args struct {
		m      int
		n      int
		guards [][]int
		walls  [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				m:      4,
				n:      6,
				guards: [][]int{{0, 0}, {1, 1}, {2, 3}},
				walls:  [][]int{{0, 1}, {2, 2}, {1, 4}},
			},
			want: 7,
		},
		{
			name: "test case 2",
			args: args{
				m:      3,
				n:      3,
				guards: [][]int{{1, 1}},
				walls:  [][]int{{0, 1}, {1, 0}, {2, 1}, {1, 2}},
			},
			want: 4,
		},
		{
			name: "test case 3",
			args: args{
				m:      5,
				n:      5,
				guards: [][]int{{1, 4}, {4, 1}, {0, 3}},
				walls:  [][]int{{3, 2}},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countUnguarded(tt.args.m, tt.args.n, tt.args.guards, tt.args.walls); got != tt.want {
				t.Errorf("countUnguarded() = %v, want %v", got, tt.want)
			}
		})
	}
}
