package main

import "fmt"

func main() {
	fmt.Println(ReverseAlphabetValue('g'))
}

func ReverseAlphabetValue(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		return 'z' - (ch - 'a')
	}

	if ch >= 'A' && ch <= 'Z' {
		return 'Z' - (ch - 'A')
	}

	return ch
}


