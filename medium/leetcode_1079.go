package main

func numTilePossibilities(tiles string) int {
	used := make(map[int]bool)
	combination := make(map[string]bool)

	permute([]rune(tiles), "", used, combination)
	return len(combination)
}

func permute(tiles []rune, output string, used map[int]bool, combination map[string]bool) {
	if len(output) > 0 {
		combination[output] = true
	}

	for i := 0; i < len(tiles); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		permute(tiles, output+string(tiles[i]), used, combination)
		used[i] = false
	}
}
