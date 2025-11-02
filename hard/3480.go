package hard

func maxSubarrays(n int, conflictingPairs [][]int) int64 {
	//totalSubarrays := n * (n + 1) / 2
	maxValidSubarrays := int64(0)

	// 각 충돌 쌍을 제거했을 때의 결과 계산
	for removeIndex := 0; removeIndex < len(conflictingPairs); removeIndex++ {
		invalidCount := 0

		// 남은 충돌 쌍들로 인해 무효가 되는 부분배열 개수 계산
		for i := 0; i < len(conflictingPairs); i++ {
			if i == removeIndex {
				continue // 제거된 쌍은 무시
			}

			pair := conflictingPairs[i]
			a, b := pair[0], pair[1]

			// a와 b를 모두 포함하는 부분배열 개수 계산
			if a > b {
				a, b = b, a // a <= b 보장
			}

			// a와 b를 모두 포함하는 부분배열 개수
			// start는 1~a, end는 b~n 범위에서 선택
			invalidForThisPair := a * (n - b + 1)
			invalidCount += invalidForThisPair
		}

		// 중복 계산된 부분 처리 (복잡하므로 간단한 방법 사용)
		validCount := countValidSubarraysSimple(n, conflictingPairs, removeIndex)

		if validCount > maxValidSubarrays {
			maxValidSubarrays = validCount
		}
	}

	return maxValidSubarrays
}

func countValidSubarraysSimple(n int, conflictingPairs [][]int, removeIndex int) int64 {
	count := int64(0)

	for start := 1; start <= n; start++ {
		for end := start; end <= n; end++ {
			valid := true

			// 제거되지 않은 충돌 쌍들만 확인
			for i, pair := range conflictingPairs {
				if i == removeIndex {
					continue
				}

				a, b := pair[0], pair[1]
				if a >= start && a <= end && b >= start && b <= end {
					valid = false
					break
				}
			}

			if valid {
				count++
			}
		}
	}

	return count
}
