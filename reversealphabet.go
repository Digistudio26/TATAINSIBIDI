package main

func ReverseAlphabet(step int) string {
	if step <= 0 {
		step = 1
	}

	result := ""
	for char := 'z'; char >= 'a'; char -= rune(step) {
		result += string(char)
	}
	return result
}

func main() {}