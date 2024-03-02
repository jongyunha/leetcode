package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_roman_to_integer(t *testing.T) {
	s := "III"
	excepted := 3
	result := romanToInt(s)
	assert.Equal(t, excepted, result)

	s = "LVIII"
	excepted = 58
	result = romanToInt(s)
	assert.Equal(t, excepted, result)

	s = "MCMXCIV"
	excepted = 1994
	result = romanToInt(s)
	assert.Equal(t, excepted, result)

	s = "DCXXI"
	excepted = 621
	result = romanToInt(s)
	assert.Equal(t, excepted, result)
}
