package main

func predictTheWinner(nums []int) bool {
	return dfsPredictTheWinner(0, 0, nums, true)
}

func dfsPredictTheWinner(p1 int, p2 int, nums []int, turn bool) bool {
	if len(nums) == 0 {
		return p1 >= p2
	}
	if len(nums) == 1 {
		if turn {
			p1 += nums[0]
		} else {
			p2 += nums[0]
		}
		return p1 >= p2
	}
	first := nums[0]
	last := nums[len(nums)-1]
	if turn {
		// Player 1은 자신이 이길 수 있는 선택이 하나라도 있으면 그 선택을 한다.
		chooseFirst := dfsPredictTheWinner(
			p1+first,
			p2,
			nums[1:],
			false,
		)
		chooseLast := dfsPredictTheWinner(
			p1+last,
			p2,
			nums[:len(nums)-1],
			false,
		)
		return chooseFirst || chooseLast
	}
	// Player 2는 Player 1을 패배시킬 수 있는 선택을 한다.
	// 따라서 Player 1은 Player 2의 두 선택 모두에 대해 이겨야 한다.
	chooseFirst := dfsPredictTheWinner(
		p1,
		p2+first,
		nums[1:],
		true,
	)
	chooseLast := dfsPredictTheWinner(
		p1,
		p2+last,
		nums[:len(nums)-1],
		true,
	)
	return chooseFirst && chooseLast
}
