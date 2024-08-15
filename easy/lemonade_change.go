package easy

func lemonadeChange(bills []int) bool {
	change5 := 0
	change10 := 0

	for _, bill := range bills {
		if bill == 5 {
			change5++
		} else if bill == 10 {
			if change5 < 1 {
				return false
			}
			change5--
			change10++
		} else if bill == 20 {
			if change10 >= 1 && change5 >= 1 {
				change5--
				change10--
			} else if change5 >= 3 {
				change5 -= 3
			} else {
				return false
			}

		}
	}

	return true
}
