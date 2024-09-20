package hard

import "math"

func shortestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	if isPalindrome(s) {
		return s
	}

	for i := len(s); i > 0; i-- {
		if isPalindrome(s[:i]) {
			return Reverse(s[i:]) + s
		}
	}

	return ""
}

func isPalindrome(s string) bool {
	for i := 0; i < int(math.Floor(float64(len(s)/2))); i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
