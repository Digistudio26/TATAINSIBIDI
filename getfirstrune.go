package sprint
func GetFirstRune(s string) rune {
	for _, r := range s {
		Println(r)
		return r
	}
	return 0
}

