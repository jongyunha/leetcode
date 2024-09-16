package main

import "testing"

func Test_minimum_time_difference_test(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want int
	}{
		{
			name: "test case 1",
			args: []string{"23:59", "00:00"},
			want: 1,
		},
		{
			name: "test case 2",
			args: []string{"00:00", "23:59", "00:00"},
			want: 0,
		},
		{
			name: "test case 3",
			args: []string{"00:00", "23:59"},
			want: 1,
		},
		{
			name: "test case 4",
			args: []string{"12:00", "00:00"},
			want: 720,
		},
		{
			name: "test case 5",
			args: []string{"01:01", "02:01"},
			want: 60,
		},
		{
			name: "test case 5",
			args: []string{"01:01", "02:01", "03:00"},
			want: 59,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinDifference(tt.args); got != tt.want {
				t.Errorf("findMinDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
