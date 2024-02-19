package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicateFromSortedArrayII(t *testing.T) {
	case1 := []int{1, 1, 1, 2, 2, 3}

	excepted := []int{1, 1, 2, 2, 3}

	output := removeDuplicates(case1)

	assert.Equal(t, output, 5)
	assert.Equal(t, case1[:output], excepted)

	case2 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	excepted2 := []int{0, 0, 1, 1, 2, 3, 3}
	output2 := removeDuplicates(case2)

	assert.Equal(t, output2, 7)
	assert.Equal(t, case2[:output2], excepted2)

}
