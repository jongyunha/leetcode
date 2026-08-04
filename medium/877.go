package main

func stoneGame(piles []int) bool {
	return dfsStoneGame(0, 0, piles, true)
}

func dfsStoneGame(p1 int, p2 int, piles []int, turn bool) bool {
	if len(piles) == 0 {
		return p1 > p2
	}

	first := piles[1:]
	last := piles[:len(piles)-1]

	if turn {
		return dfsStoneGame(p1+piles[0], p2, first, false) || dfsStoneGame(p1+piles[len(piles)-1], p2, last, false)
	}

	return dfsStoneGame(p1, p2+piles[0], first, true) && dfsStoneGame(p1, p2+piles[len(piles)-1], last, true)
}
