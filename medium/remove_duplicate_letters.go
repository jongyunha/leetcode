package main

// Given a string s, remove duplicate letters so that every letter appears once and only once. You must make sure your result is
// the smallest in lexicographical order
// among all possible results.
//
// Example 1:
//
// Input: s = "bcabc"
// Output: "abc"
// Example 2:
//
// Input: s = "cbacdcbc"
// Output: "acdb"
func removeDuplicateLetters(s string) string {
	lastSequence := make(map[byte]int)
	seen := make(map[byte]struct{})
	stack := make([]byte, 0)

	for i := range s {
		lastSequence[s[i]] = i
	}

	for i := range s {
		if _, ok := seen[s[i]]; !ok {
			for len(stack) > 0 && s[i] < stack[len(stack)-1] && lastSequence[stack[len(stack)-1]] > i {
				delete(seen, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, s[i])
			seen[s[i]] = struct{}{}
		}
	}

	return string(stack)
}
