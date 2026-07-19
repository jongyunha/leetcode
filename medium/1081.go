package main

func smallestSubsequence(s string) string {
	n := len(s)
	stack := make([]int, 0)
	lastIndex := make([]int, 26)
	for i := 0; i < n; i++ {
		lastIndex[s[i]-'a'] = i
	}

	visited := make(map[byte]bool)
	for i := range s {
		curr := s[i]
		if visited[curr] {
			continue
		}

		for len(stack) > 0 {
			topIndex := stack[len(stack)-1]
			top := s[topIndex]
			if curr >= top {
				break
			}
			if lastIndex[top-'a'] <= i {
				break
			}
			visited[top] = false
			stack = stack[:len(stack)-1]
		}
		visited[curr] = true
		stack = append(stack, i)
	}

	answer := make([]byte, 0, len(stack))
	for _, index := range stack {
		answer = append(answer, s[index])
	}
	return string(answer)
}
