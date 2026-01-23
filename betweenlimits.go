package main 

import "strings"

func BetweenLimits(from, to rune) string {
	// Ensure correct order
	if from > to {
		from, to = to, from
	}

	var b strings.Builder

	// Start after `from`, stop before `to`
	for r := from + 1; r < to; r++ {
		b.WriteRune(r)
	}

	return b.String()
}
