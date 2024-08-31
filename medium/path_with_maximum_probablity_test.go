package main

import "testing"

func Test_path_with_maximum_probablity(t *testing.T) {
	type args struct {
		n          int
		edges      [][]int
		succProb   []float64
		start_node int
		end_node   int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test-case-1",
			args: args{
				n: 3,
				edges: [][]int{
					{0, 1},
					{1, 2},
					{0, 2},
				},
				succProb:   []float64{0.5, 0.5, 0.2},
				start_node: 0,
				end_node:   2,
			},
			want: 0.25,
		},
		{
			name: "test-case-2",
			args: args{
				n: 3,
				edges: [][]int{
					{0, 1},
					{1, 2},
					{0, 2},
				},
				succProb:   []float64{0.5, 0.5, 0.3},
				start_node: 0,
				end_node:   2,
			},
			want: 0.3,
		},
		{
			name: "test-case-3",
			args: args{
				n: 3,
				edges: [][]int{
					{0, 1},
					{1, 2},
					{0, 2},
				},
				succProb:   []float64{0.5, 0.5, 0.3},
				start_node: 0,
				end_node:   1,
			},
			want: 0.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProbability(tt.args.n, tt.args.edges, tt.args.succProb, tt.args.start_node, tt.args.end_node); got != tt.want {
				t.Errorf("maxProbability() = %v, want %v", got, tt.want)
			}
		})
	}
}
