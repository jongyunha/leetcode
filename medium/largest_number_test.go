package main

import "testing"

func Test_largest_number(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Case 1", args{[]int{10, 2}}, "210"},
		{"Case 2", args{[]int{3, 30, 34, 5, 9}}, "9534330"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestNumber(tt.args.nums); got != tt.want {
				t.Errorf("largest_number() = %v, want %v", got, tt.want)
			}
		})
	}
}
