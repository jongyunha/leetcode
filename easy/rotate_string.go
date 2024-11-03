package easy

import "fmt"

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	for i := range goal {
		front := goal[:i]
		rear := goal[i:]
		combine := fmt.Sprintf("%s%s", rear, front)
		if combine == s {
			return true
		}
	}

	return false
}
