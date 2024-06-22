package main

import "testing"

func Test_count_number_of_nice_subarrays(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test #1", args{[]int{1, 1, 2, 1, 1}, 3}, 2},
		{"test #2", args{[]int{2, 4, 6}, 1}, 0},
		{"test #3", args{[]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSubarrays(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("count_number_of_nice_subarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
