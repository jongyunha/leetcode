package easy

import "testing"

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	var max int
	var maxNum int
	for num, count := range m {
		if count > max {
			max = count
			maxNum = num
		}
	}

	return maxNum
}


func TestMajorityElement(t *testing.T) {
	nums := []int{3, 2, 3}
	result := majorityElement(nums)
	if result != 3 {
		t.Error("Wrong result")
	}
}