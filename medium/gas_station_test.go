package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_gas_station(t *testing.T) {

	//gas := []int{1, 2, 3, 4, 5}
	//costs := []int{3, 4, 5, 1, 2}
	//result := canCompleteCircuit(gas, costs)
	//excepted := 3
	//
	//assert.Equal(t, excepted, result)
	//
	gas2 := []int{2, 3, 4}
	costs2 := []int{3, 4, 3}
	result2 := canCompleteCircuit(gas2, costs2)
	excepted2 := -1
	assert.Equal(t, excepted2, result2)
}
