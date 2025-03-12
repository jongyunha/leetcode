package main

func numberOfSubstrings(s string) int {
	result := 0
	n := len(s)

	// 각 문자의 개수를 저장할 맵
	count := map[byte]int{
		'a': 0,
		'b': 0,
		'c': 0,
	}

	left := 0 // 윈도우의 왼쪽 경계

	for right := 0; right < n; right++ {
		// 오른쪽 문자 추가
		count[s[right]]++

		// 'a', 'b', 'c'가 모두 최소 한 번씩 포함될 때까지 윈도우 확장
		for left <= right && count['a'] > 0 && count['b'] > 0 && count['c'] > 0 {
			// 현재 윈도우로부터 만들 수 있는 유효한 부분 문자열의
			// 개수는 n - right임 (right부터 끝까지 어디서든 끝날 수 있음)
			result += n - right

			// 왼쪽 경계 이동
			count[s[left]]--
			left++
		}
	}

	return result
}
