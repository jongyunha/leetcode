package main

import "strconv"

func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}

	if len(tokens) == 1 {
		i, _ := strconv.Atoi(tokens[0])
		return i
	}

	intStack := make([]int, 0)

	for _, s := range tokens {
		switch s {
		case "+":
			i1, i2 := intStack[len(intStack)-1], intStack[len(intStack)-2]
			intStack = intStack[:len(intStack)-2]
			r := i1 + i2
			intStack = append(intStack, r)
		case "*":
			i1, i2 := intStack[len(intStack)-1], intStack[len(intStack)-2]
			intStack = intStack[:len(intStack)-2]
			r := i1 * i2
			intStack = append(intStack, r)
		case "-":
			i1, i2 := intStack[len(intStack)-1], intStack[len(intStack)-2]
			intStack = intStack[:len(intStack)-2]
			r := i2 - i1
			intStack = append(intStack, r)
		case "/":
			i1, i2 := intStack[len(intStack)-1], intStack[len(intStack)-2]
			intStack = intStack[:len(intStack)-2]
			r := i2 / i1
			intStack = append(intStack, r)
		default:
			i, _ := strconv.Atoi(s)
			intStack = append(intStack, i)
		}
	}

	return intStack[0]
}
