package main

func smallestPalindrome(s string) string {
	counter := make([]int, 26)
	for _, c := range s {
		counter[c-'a']++
	}

	i := 0
	reverse := false
	answer := make([]byte, len(s))
	filed := 0
	answerIdx := 0
	n := len(s)
	if n%2 != 0 {
		filed++
		var middle rune
		for j := 25; j >= 0; j-- {
			if counter[j] > 0 && counter[j]%2 != 0 {
				middle = rune(j + 'a')
				counter[j]--
			}
		}
		answer[n/2] = byte(middle)
	}

	for filed != len(s) {
		if counter[i] == 0 {
			i++
			continue
		}
		counter[i]--
		char := i + 'a'
		if !reverse {
			answer[answerIdx] = byte(char)
			reverse = true
			answerIdx++
		} else {
			answer[len(answer)-answerIdx] = byte(char)
			reverse = false
		}
		filed++
	}
	return string(answer)
}
