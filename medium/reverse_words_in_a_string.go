package main

import (
	"strings"
)

func reverseWords(s string) string {
	removeTrimed := strings.Trim(s, " ")
	sb := strings.Builder{}
	queue := make([]string, 0)

	var prev rune
	for _, c := range removeTrimed {
		if c != ' ' {
			sb.WriteRune(c)
		} else if c == ' ' && prev != ' ' {
			queue = append(queue, sb.String())
			sb.Reset()
		}

		prev = c
	}

	queue = append(queue, sb.String())
	sb.Reset()

	for len(queue) > 1 {
		pop := queue[len(queue)-1]
		queue = queue[0 : len(queue)-1]

		sb.WriteString(pop)
		sb.WriteString(" ")
	}

	pop := queue[len(queue)-1]
	sb.WriteString(pop)

	return sb.String()
}
