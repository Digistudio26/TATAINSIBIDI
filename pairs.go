package main

import "fmt"

func Pairs() string {
	result := ""
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			pair := fmt.Sprintf("%02d %02d", i, j)
			if result != "" {
				result += ","
			}
			result += pair
		}
	}
	return result
}

func main() {
	fmt.Println(Pairs())
}
