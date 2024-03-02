package easy

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for idx, num := range nums {
		m[num] = idx
	}

	for idx, num := range nums {
		if compareIdx, ok := m[target-num]; ok && compareIdx != idx {
			return []int{idx, compareIdx}
		}
	}

	return nil
}

func main() {
	nums := []int{2, 7, 5, 9, 4}
	result := twoSum(nums, 9)
	fmt.Println(result)
}
