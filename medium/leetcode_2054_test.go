package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxTwoEvents(t *testing.T) {
	event1 := make([][]int, 3)
	event1[0] = []int{1, 3, 2}
	event1[1] = []int{4, 5, 2}
	event1[2] = []int{2, 4, 3}

	event2 := make([][]int, 3)
	event2[0] = []int{1, 3, 2}
	event2[1] = []int{4, 5, 2}
	event2[2] = []int{1, 5, 5}

	// event3 [[66,97,90],[98,98,68],[38,49,63],[91,100,42],[92,100,22],[1,77,50],[64,72,97]]
	event3 := make([][]int, 7)
	event3[0] = []int{66, 97, 90}
	event3[1] = []int{98, 98, 68}
	event3[2] = []int{38, 49, 63}
	event3[3] = []int{91, 100, 42}
	event3[4] = []int{92, 100, 22}
	event3[5] = []int{1, 77, 50}
	event3[6] = []int{64, 72, 97}

	tests := []struct {
		events [][]int
		want   int
	}{
		// {
		// 	events: event2,
		// 	want:   5,
		// },
		// {
		// 	events: event1,
		// 	want:   4,
		// },
		{
			events: event3,
			want:   165,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, maxTwoEvents(tt.events))
	}
}
