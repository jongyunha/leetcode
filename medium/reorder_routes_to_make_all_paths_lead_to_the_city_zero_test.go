package main

import "testing"

func Test_reorder_routes_to_make_all_paths_lead_to_the_city_zero(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		connections [][]int
		want        int
	}{
		{
			name: "test1",
			n:    6,
			connections: [][]int{
				{0, 1},
				{1, 3},
				{2, 3},
				{4, 0},
				{4, 5},
			},
			want: 3,
		},
		{
			name: "test2",
			n:    5,
			connections: [][]int{
				{1, 0},
				{1, 2},
				{3, 2},
				{3, 4},
			},
			want: 2,
		},
		{
			name: "test3",
			n:    3,
			connections: [][]int{
				{1, 0},
				{2, 0},
			},
			want: 0,
		},
		{
			name: "test4",
			n:    5,
			connections: [][]int{
				{1, 0},
				{1, 2},
				{3, 2},
				{3, 4},
				{4, 2},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minReorder(tt.n, tt.connections)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
