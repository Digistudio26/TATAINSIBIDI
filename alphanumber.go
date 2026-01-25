
package main




func AlphaNumber(n int) string {
	if n == 0 {
		return "a"
	}

	result := ""
	sign := ""

	if n < 0 {
		sign = "-"
		n = -n
	}

	for n > 0 {
		digit := n % 10
		letter := rune('a' + digit)
		result = string(letter) + result
		n /= 10
	}

	return sign + result
}
