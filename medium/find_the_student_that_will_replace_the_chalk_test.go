package main

import "testing"

func Test_find_the_student_that_will_replace_the_chalk(t *testing.T) {
	type args struct {
		chalk []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test_case_1", args{[]int{5, 1, 5}, 22}, 0},
		{"test_case_2", args{[]int{3, 4, 1, 2}, 25}, 1},
		{"test_case_4", args{[]int{3, 22, 9, 28, 48, 40, 47, 5, 19, 4, 9, 7, 11, 48}, 953}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chalkReplacer(tt.args.chalk, tt.args.k); got != tt.want {
				t.Errorf("chalkReplacer() = %v, want %v", got, tt.want)
			}
		})
	}
}
