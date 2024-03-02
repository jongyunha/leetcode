package main

func canCompleteCircuit(gas []int, cost []int) int {
	totalGas, totalCost, start, tank := 0, 0, 0, 0

	for i := 0; i < len(gas); i++ {
		currGas := gas[i]
		currCost := cost[i]

		totalGas += currGas
		totalCost += currCost
		tank += currGas - currCost
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}

	if totalCost > totalGas {
		return -1
	}

	return start
}
