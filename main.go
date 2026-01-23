package main

import "fmt"
import "math"

func Casting(n float64) int {
	return int(math.Round(n))
}

func main() { 

	fmt.Println(Casting(1.85))
}