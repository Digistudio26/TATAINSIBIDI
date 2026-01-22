package main

import "fmt"

func main() {
	fmt.Println(IntVsFloat(5, 8.8))
}

func IntVsFloat(i int, f float32) string {
	if float32(i) > f {
		return "Integer"
	} else if float32(i) < f {
		return "Float"
	}
	return "Same"
}
