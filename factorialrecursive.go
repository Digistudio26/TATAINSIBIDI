package main


func FactorialIterative(n int) int {
	if n < 0 {
		return 0 // factorial not defined for negatives
	}

	result := 1
	for i := 1; i <= n; i++ { // multiply 1 × 2 × ... × n
		result *= i
	}

	return result
}
