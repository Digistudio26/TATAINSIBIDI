package main

func FactorialIterative(n int) int {
	if n < 0 {
		return 0
	}

	result := 1

	for i := n; i >= 1; i-- {
		result *= i
	}

	return result
}
