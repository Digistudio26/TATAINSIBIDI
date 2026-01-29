package main


func ToCapitalCase(s string) string {
	runes := []rune(s)
	newWord := true

	for i, r := range runes {
		// Check if alphanumeric
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			if newWord {
				// Capitalize first letter of the word
				if r >= 'a' && r <= 'z' {
					runes[i] = r - 32
				}
				newWord = false
			} else {
				// Lowercase the rest of the word
				if r >= 'A' && r <= 'Z' {
					runes[i] = r + 32
				}
			}
		} else {
			// Non-alphanumeric character → next character starts a new word
			newWord = true
		}
	}
	return string(runes)
}
 