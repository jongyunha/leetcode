package main

func singleNumber(nums []int) int {
	uniqueMap := make(map[int]int)

	for _, num := range nums {
		uniqueMap[num]++
	}

	found := 0

	for key := range uniqueMap {
		if uniqueMap[key] == 1 {
			found = key
		}
	}

	return found
}
