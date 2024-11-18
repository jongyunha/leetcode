package easy

func decrypt(code []int, k int) []int {
	n := len(code)
	decode := make([]int, len(code))

	for i := 0; i < n; i++ {
		if k > 0 {
			sum := 0
			for j := i + 1; j <= k+i; j++ {
				index := (j + n) % n
				sum += code[index]
			}
			decode[i] = sum
		} else if k < 0 {
			sum := 0
			for j := 0; j < -k; j++ {
				index := (n - j - 1 + i) % n
				sum += code[index]
			}
			decode[i] = sum
		} else {
			decode[i] = 0
		}
	}

	return decode
}
