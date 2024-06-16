package main

import (
	"math"
)

func maxVowels(s string, k int) int {
	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	continuous := 0

	i, j := 0, 0

	localContinuous := 0

	for j < len(s) {
		if j < k {
			c := rune(s[j])
			if _, ok := vowels[c]; ok {
				localContinuous++
			}
			j++
			continuous = int(math.Max(float64(localContinuous), float64(continuous)))
		} else {
			prev := rune(s[i])
			if _, ok := vowels[prev]; ok {
				localContinuous--
			}
			i++
			next := rune(s[j])
			if _, ok := vowels[next]; ok {
				localContinuous++
			}
			j++
			continuous = int(math.Max(float64(localContinuous), float64(continuous)))
		}
	}
	return continuous
}
