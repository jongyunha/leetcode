package easy

import "testing"

// Given a 2D grid of size m x n and an integer k. You need to shift the grid k times.
//
// In one shift operation:
//
// Element at grid[i][j] moves to grid[i][j + 1].
// Element at grid[i][n - 1] moves to grid[i + 1][0].
// Element at grid[m - 1][n - 1] moves to grid[0][0].
// Return the 2D grid after applying shift operation k times.
//
// Example 1:
//
// Input: grid = [[1,2,3],[4,5,6],[7,8,9]], k = 1
// Output: [[9,1,2],[3,4,5],[6,7,8]]
// Example 2:
//
// Input: grid = [[3,8,1,9],[19,7,2,5],[4,6,11,10],[12,0,21,13]], k = 4
// Output: [[12,0,21,13],[3,8,1,9],[19,7,2,5],[4,6,11,10]]
// Example 3:
//
// Input: grid = [[1,2,3],[4,5,6],[7,8,9]], k = 9
// Output: [[1,2,3],[4,5,6],[7,8,9]]
func TestShiftGrid(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		k    int
		want [][]int
	}{
		{
			name: "test1",
			grid: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			k:    1,
			want: [][]int{{9, 1, 2}, {3, 4, 5}, {6, 7, 8}},
		},
		{
			name: "test2",
			grid: [][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}},
			k:    4,
			want: [][]int{{12, 0, 21, 13}, {3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}},
		},
		{
			name: "test3",
			grid: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			k:    9,
			want: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shiftGrid(tt.grid, tt.k); !equal2D(got, tt.want) {
				t.Errorf("shiftGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal2D(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
