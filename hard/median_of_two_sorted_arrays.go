package hard

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := merge(nums1, nums2)
	var result float64
	if len(merged)%2 == 0 {
		result = float64(merged[len(merged)/2-1]+merged[len(merged)/2]) / 2.0
	} else {
		result = float64(merged[len(merged)/2])
	}

	return result
}

func merge(a []int, b []int) []int {
	final := make([]int, 0, len(a)+len(b))
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	for i < len(a) {
		final = append(final, a[i])
		i++
	}

	for j < len(b) {
		final = append(final, b[j])
		j++
	}

	return final
}
