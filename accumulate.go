package main

import "fmt"

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

func main() {
	fmt.Println(Accumulate(4))  // Output: 10
	fmt.Println(Accumulate(0))  // Output: 0
	fmt.Println(Accumulate(-5)) // Output: 0
}
