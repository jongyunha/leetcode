package main

import "testing"

func Test_house_robber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "test 1",
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
		{
			name: "test 2",
			nums: []int{2, 7, 9, 3, 1},
			want: 12,
		},
		{
			name: "test 3",
			nums: []int{2, 1, 1, 2},
			want: 4,
		},
		{
			name: "test 4",
			nums: []int{1, 2, 3, 1, 1, 100},
			want: 104,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.nums); got != tt.want {
				t.Errorf("house_robber() = %v, want %v", got, tt.want)
			}
		})
	}
}
