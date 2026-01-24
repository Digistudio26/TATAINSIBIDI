package main

import "fmt"


func Combinations() string {
	result := ""
	first := true

	for i := 0; i <= 9; i++ {
		for j := i + 1; j <= 9; j++ {
			for k := j + 1; k <= 9; k++ {
				if !first {
					result += ", "
				}
				result += fmt.Sprintf("%d%d%d", i, j, k)
				first = false
			}
		}
	}

	return result
}
