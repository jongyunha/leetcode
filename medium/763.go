package main

func partitionLabels(s string) []int {
	lastIndexMap := make(map[rune]int)

	for i, c := range s {
		lastIndexMap[c] = i
	}

	result := make([]int, 0)

	start := s[0]
	accumulate := -1
	lastIndex := lastIndexMap[rune(start)]

	for i := 1; i < len(s); i++ {
		if i <= lastIndex {
			lastIndex = max(lastIndex, lastIndexMap[rune(s[i])])
		} else {
			result = append(result, lastIndex-accumulate)
			accumulate = lastIndex
			lastIndex = lastIndexMap[rune(s[i])]
		}
	}
	result = append(result, lastIndex-accumulate)
	return result
}
