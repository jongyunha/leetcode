package main

import (
	"sort"
	"strings"
)

func clearStars(s string) string {
	stacks := make([][]int, 26) // 'a' ~ 'z'

	for i, char := range s {
		if char == '*' {
			for j := 0; j < 26; j++ {
				if len(stacks[j]) > 0 {
					stacks[j] = stacks[j][:len(stacks[j])-1]
					break
				}
			}
		} else {
			stacks[char-'a'] = append(stacks[char-'a'], i)
		}
	}

	var remainingIndices []int
	for i := 0; i < 26; i++ {
		for _, index := range stacks[i] {
			remainingIndices = append(remainingIndices, index)
		}
	}

	sort.Ints(remainingIndices)

	var result strings.Builder
	result.Grow(len(remainingIndices))
	for _, index := range remainingIndices {
		result.WriteByte(s[index])
	}

	return result.String()
}
