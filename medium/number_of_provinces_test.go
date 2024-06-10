package main

import "testing"

func Test_number_of_provinces(t *testing.T) {
	tests := []struct {
		name        string
		isConnected [][]int
		want        int
	}{
		{
			name: "test1",
			isConnected: [][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 1},
			},
			want: 2,
		},
		{
			name: "test2",
			isConnected: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			want: 3,
		},
		{
			name: "test3",
			isConnected: [][]int{
				{1, 0, 0, 1},
				{0, 1, 1, 0},
				{0, 1, 1, 0},
				{1, 0, 0, 1},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findCircleNum(tt.isConnected)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
