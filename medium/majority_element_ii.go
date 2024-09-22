package main

func majorityElement(nums []int) []int {
	// Moore's Voting Algorithm
	// We need at most two majority elements

	// Variables to keep track of the count of the candidates
	count1, count2 := 0, 0
	// Variables for the majority element candidates
	var majority1, majority2 int
	// Boolean flags to track if the majority elements have been initialized
	init1, init2 := false, false

	// First pass to find candidates for majority elements
	for _, num := range nums {
		// If the current number is the majority element 1 candidate
		if init1 && num == majority1 {
			count1++
		} else if init2 && num == majority2 {
			count2++
		} else if count1 == 0 {
			majority1 = num
			count1 = 1
			init1 = true
		} else if count2 == 0 {
			majority2 = num
			count2 = 1
			init2 = true
		} else {
			count1--
			count2--
		}
	}

	// Reset the counts
	count1, count2 = 0, 0

	// Second pass to confirm the majority elements
	for _, num := range nums {
		if init1 && num == majority1 {
			count1++
		}
		if init2 && num == majority2 {
			count2++
		}
	}

	// Initialize a slice to store the majority elements
	var majorityElements []int

	// Check if the candidates appear more than n/3 times
	if count1 > len(nums)/3 {
		majorityElements = append(majorityElements, majority1)
	}
	if count2 > len(nums)/3 {
		majorityElements = append(majorityElements, majority2)
	}

	// Return the list of majority elements
	return majorityElements
}
