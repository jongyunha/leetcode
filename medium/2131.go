package main

func longestPalindrome(words []string) int {
	wordCounts := make(map[string]int)

	for _, word := range words {
		wordCounts[word]++
	}

	length := 0
	hasOddSelfPalindrome := false

	for word, count := range wordCounts {
		if count == 0 {
			continue
		}

		reverseWord := string(word[1]) + string(word[0])
		reverseCount := wordCounts[reverseWord]

		if word == reverseWord {
			length += (count / 2) * 4
			if count%2 == 1 {
				hasOddSelfPalindrome = true
			}
		} else {
			pairs := min(count, reverseCount)
			length += pairs * 4
			wordCounts[word] -= pairs
			wordCounts[reverseWord] -= pairs
		}
	}

	if hasOddSelfPalindrome {
		length += 2
	}

	return length
}
