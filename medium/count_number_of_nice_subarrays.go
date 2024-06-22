package main

func numberOfSubarrays(nums []int, k int) int {
	odd := []int{}
	for i, v := range nums {
		if v%2 == 1 {
			odd = append(odd, i)
		}
	}
	if len(odd) < k {
		return 0
	}

	res := 0
	for i := 0; i < len(odd)-k+1; i++ {
		left, right := 0, 0

		if i == 0 {
			left = odd[i] + 1
		} else {
			left = odd[i] - odd[i-1]
		}

		if i+k == len(odd) {
			right = len(nums) - odd[i+k-1]
		} else {
			right = odd[i+k] - odd[i+k-1]
		}

		res += left * right
	}

	return res
}
