package easy

func isOneBitCharacter(bits []int) bool {
	if len(bits) == 0 {
		return false
	}

	if len(bits) == 1 && bits[0] == 0 {
		return true
	}

	cover := 0
	for i := 0; i < len(bits)-1; {
		a := bits[i]
		if a == 0 {
			cover = i
			i++
		} else {
			cover = i + 1
			i += 2
		}
	}
	if cover == len(bits)-1 {
		return false
	}
	return true
}
