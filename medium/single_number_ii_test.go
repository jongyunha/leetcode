package main

import "testing"

func Test_single_number_ii(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 2, 3, 2}, 3},
		{[]int{0, 1, 0, 1, 0, 1, 99}, 99},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := singleNumber(tt.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
