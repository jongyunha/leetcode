package main

import "testing"

func TestTriangleNumber(t *testing.T) {
	type args struct {
		nums   []int
		output int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Example 1",
			args: args{
				nums:   []int{2, 2, 3, 4},
				output: 3,
			},
		},
		{
			name: "Example 2",
			args: args{
				nums:   []int{4, 2, 3, 4},
				output: 4,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangleNumber(tt.args.nums); got != tt.args.output {
				t.Errorf("triangleNumber() = %v, want %v", got, tt.args.output)
			}
		})
	}
}
