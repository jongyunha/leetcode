package main

import "testing"

func Test_magnetic_force_between_two_balls(t *testing.T) {
	type args struct {
		position []int
		m        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{[]int{1, 2, 3, 4, 7}, 3}, 3},
		{"test 2", args{[]int{5, 4, 3, 2, 1, 1000000000}, 2}, 999999999},
		{"test 3", args{[]int{79, 74, 57, 22}, 4}, 5},
		{"test 4", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistance(tt.args.position, tt.args.m); got != tt.want {
				t.Errorf("maxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
