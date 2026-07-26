package main

func leastInterval(tasks []byte, n int) int {
	taskFrequency := make([]int, 26)
	for _, task := range tasks {
		taskFrequency[task-'A']++
	}

	maxFrequency := 0
	maxFrequencyCount := 0

	for _, frequency := range taskFrequency {
		if frequency > maxFrequency {
			maxFrequencyCount = 1
			maxFrequency = frequency
		} else if frequency == maxFrequency {
			maxFrequencyCount++
		}
	}

	requiredInterval := (maxFrequency-1)*(n+1) + maxFrequencyCount
	if requiredInterval > len(tasks) {
		return requiredInterval
	}

	return len(tasks)
}
