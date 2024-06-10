package main

import "testing"

func Test_keys_and_rooms(t *testing.T) {
	type args struct {
		rooms [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testcase 1",
			args: args{
				rooms: [][]int{
					{1},
					{2},
					{3},
					{},
				},
			},
			want: true,
		},
		{
			name: "testcase 2",
			args: args{
				rooms: [][]int{
					{1, 3},
					{3, 0, 1},
					{2},
					{0},
				},
			},
			want: false,
		},
		{
			name: "testcase 3",
			args: args{
				rooms: [][]int{
					{2},
					{},
					{1},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canVisitAllRooms(tt.args.rooms); got != tt.want {
				t.Errorf("keys_and_rooms() = %v, want %v", got, tt.want)
			}
		})
	}
}
