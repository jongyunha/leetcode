package main

func checkDivisibility(n int) bool {
	sum := 0
	product := 1
	origin := n

	for n > 0 {
		sum += n % 10
		product *= n % 10
		n /= 10
	}

	return origin%(sum+product) == 0
}
