package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_integer_to_roman_test(t *testing.T) {
	num := 1994
	excepted := "MCMXCIV"
	result := intToRoman(num)

	assert.Equal(t, excepted, result)

	num = 58
	excepted = "LVIII"
	result = intToRoman(num)

	assert.Equal(t, excepted, result)
}
