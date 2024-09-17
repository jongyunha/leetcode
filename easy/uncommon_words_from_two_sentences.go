package easy

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	split1 := strings.Split(s1, " ")
	split2 := strings.Split(s2, " ")

	wordMap := make(map[string]int)

	for _, s := range split1 {
		wordMap[s]++
	}

	for _, s := range split2 {
		wordMap[s]++
	}

	commons := make([]string, 0)

	for key, value := range wordMap {
		if value == 1 {
			commons = append(commons, key)
		}
	}

	return commons
}
