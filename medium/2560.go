package main

func minCapability(nums []int, k int) int {
	left, right := nums[0], nums[0]
	for _, num := range nums {
		if num < left {
			left = num
		}
		if num > right {
			right = num
		}
	}

	// 이진 탐색 시작
	for left < right {
		mid := (left + right) / 2
		if isPossible(nums, k, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

// 현재 capability로 k개 이상 집을 훔칠 수 있는지 판단하는 함수
func isPossible(nums []int, k int, capability int) bool {
	count, i := 0, 0
	for i < len(nums) {
		if nums[i] <= capability {
			count++
			i += 2 // 인접한 집을 건너뜀
		} else {
			i++
		}
		if count >= k {
			return true
		}
	}
	return false
}
