package main



func IsNumeric(s string) bool {
	for _, n := range s {
		if n < '0' || n > '9' {
			return false
		}
	}
	return true
}
