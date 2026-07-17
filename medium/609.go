package main

import "strings"

func findDuplicate(paths []string) [][]string {
	contentMap := make(map[string][]string)

	for _, path := range paths {
		split := strings.Split(path, " ")
		root := split[0]
		for _, file := range split[1:] {
			content := parseContent(file)

			if _, ok := contentMap[content]; !ok {
				contentMap[content] = make([]string, 0)
			}
			removed := removeContent(file)
			contentMap[content] = append(contentMap[content], strings.Join([]string{root, removed}, "/"))
		}
	}

	result := make([][]string, 0)
	for k, v := range contentMap {
		if len(contentMap[k]) > 1 {
			result = append(result, v)
		}
	}

	return result
}

func parseContent(s string) string {
	start := strings.Index(s, "(")
	end := strings.Index(s, ")")
	return s[start+1 : end]
}

func removeContent(s string) string {
	start := strings.Index(s, "(")
	return s[:start]
}
