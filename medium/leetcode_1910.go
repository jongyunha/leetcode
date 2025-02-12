package main

import "strings"

func removeOccurrences(s string, part string) string {
	left := 0
	n := len(part)

	for left <= len(s)-n {
		temp := s[left : left+n]
		if temp == part {
			sb := strings.Builder{}
			sb.Write([]byte(s[:left]))
			sb.Write([]byte(s[left+n:]))
			s = sb.String()
			left = 0
		} else {
			left++
		}
	}

	return s
}
