package main

func taskSchedulerII(tasks []int, space int) int64 {
	taskDay := make(map[int]int64)

	day := int64(0)
	for _, task := range tasks {
		prevDay, ok := taskDay[task]
		if !ok {
			day++
			taskDay[task] = day
			continue
		}

		if prevDay+int64(space) > day {
			day += prevDay + int64(space) - day
		}
		day++
		taskDay[task] = day
	}

	return day
}
