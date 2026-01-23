package sprint

import "fmt"

func BetweenLimits(from, to rune) string {
	// Ensure correct order
	if from > to {
		from, to = to, from
	}

	var b string

	// Start after `from`, stop before `to`
	for r := from + 1; r < to; r++ {
		 b = b + string(rune(r))
	}

	return b
}
func main() {
	fmt.Println(BetweenLimits('j', 'f'))
}
