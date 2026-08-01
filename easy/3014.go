package easy

func minimumPushes(word string) int {
	capacity := make([]int, 8)
	dial := make(map[rune]int)
	answer := 0

	level := 0
	for _, c := range word {
		if _, ok := dial[c]; !ok {
			found := 0
			for i := 0; i < len(capacity); i++ {
				if level == capacity[i] {
					capacity[i]++
					found = capacity[i]
					break
				}
			}
			dial[c] = found
			answer += dial[c]
			if capacity[len(capacity)-1] == level+1 {
				level++
			}
		} else {
			answer += dial[c]
		}
	}

	return answer
}
