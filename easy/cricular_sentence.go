package easy

import "strings"

func isCircularSentence(sentence string) bool {
	split := strings.Split(sentence, " ")
	for i := 0; i < len(split)-1; i++ {
		last := split[i][len(split[i])-1]
		nextFirst := split[i+1][0]

		if last != nextFirst {
			return false
		}
	}

	lastString := split[len(split)-1]
	if lastString[len(lastString)-1] != split[0][0] {
		return false
	}

	return true
}
