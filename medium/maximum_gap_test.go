package main

import "testing"

func Test_maximum_gap(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 6, 9, 1}, 3},
		{[]int{10}, 0},
		{[]int{1, 10000000}, 9999999},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maximumGap(tt.nums); got != tt.want {
				t.Errorf("maximumGap() = %v, want %v", got, tt.want)
			}
		})
	}
}
