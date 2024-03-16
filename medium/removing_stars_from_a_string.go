package main

func remove(slice []rune, s int) []rune {
	return append(slice[:s], slice[s+1:]...)
}

func removeStars(s string) string {
	runes := []rune(s)
	n := len(runes)
	queue := make([]rune, 0)
	for i := 0; i < n; i++ {
		curr := runes[i]
		if curr != '*' {
			queue = append(queue, curr)
		} else {
			queue = remove(queue, len(queue)-1)
		}
	}
	return string(queue)
}
