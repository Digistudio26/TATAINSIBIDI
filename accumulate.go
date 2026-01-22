package main

import "fmt"

// Accumulate returns the sum of all numbers from 0 to n if n is positive.
// If n is negative, it returns 0.
func Accumulate(n int) int {
	if n < 0 {
		return 0
	}

	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}

	return sum
}
