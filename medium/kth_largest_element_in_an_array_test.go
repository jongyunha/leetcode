package main

import "testing"

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		result int
	}{
		{
			[]int{3, 2, 1, 5, 6, 4},
			2,
			5,
		},
		{
			[]int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			4,
			4,
		},
	}

	for _, test := range tests {
		actual := findKthLargest(test.nums, test.k)
		if actual != test.result {
			t.Errorf("findKthLargest expected %d, got %d", test.result, actual)
		}
	}
}
