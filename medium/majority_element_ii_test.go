package main

import "testing"

func Test_majority_element_ii(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "test 1",
			nums: []int{3, 2, 3},
			want: []int{3},
		},
		{
			name: "test 2",
			nums: []int{1},
			want: []int{1},
		},
		{
			name: "test 3",
			nums: []int{1, 2},
			want: []int{1, 2},
		},
		{
			name: "test 5",
			nums: []int{1, 1, 1, 3, 3, 2, 2, 2},
			want: []int{1, 2},
		},
		{
			name: "test 6",
			nums: []int{1, 1, 1, 3, 3, 2, 2, 2, 4, 4, 4},
			want: []int{1, 2, 4},
		},
		{
			name: "test 7",
			nums: []int{1, 2, 3},
			want: []int{},
		},
		{
			name: "test 8",
			nums: []int{0, 0, 0},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.nums); !equal(got, tt.want) {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
