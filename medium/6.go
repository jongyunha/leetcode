package main

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	rows := make([]strings.Builder, numRows)

	currentRow := 0
	direction := 1

	for _, char := range s {
		rows[currentRow].WriteRune(char)

		if currentRow == 0 {
			direction = 1
		} else if currentRow == numRows-1 {
			direction = -1
		}

		currentRow += direction
	}

	var result strings.Builder
	result.Grow(len(s))

	for i := range rows {
		result.WriteString(rows[i].String())
	}

	return result.String()
}
