package easy

import "testing"

func Test_convert_1d_array_int_2d_array(t *testing.T) {
	type args struct {
		original []int
		m        int
		n        int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test Case 1",
			args: args{
				original: []int{1, 2, 3, 4},
				m:        2,
				n:        2,
			},
			want: [][]int{
				{1, 2},
				{3, 4},
			},
		},
		{
			name: "Test Case 1",
			args: args{
				original: []int{1, 2},
				m:        1,
				n:        1,
			},
			want: nil,
		},
		{
			name: "Test Case 3",
			args: args{
				original: []int{3},
				m:        1,
				n:        2,
			},
			want: [][]int{{
				3,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := construct2DArray(tt.args.original, tt.args.m, tt.args.n); !compare2DArray(got, tt.want) {
				t.Errorf("convert_1d_array_int_2d_array() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compare2DArray(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}

		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}
