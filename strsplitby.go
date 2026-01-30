package main


func strSplitBy(s, sep string) []string {
	if sep == "" {
		return []string{s}
	}

	var result []string
	start := 0
	i := 0

	for i <= len(s)-len(sep) {
		if s[i:i+len(sep)] == sep {
			result = append(result, s[start:i])
			i += len(sep)
			start = i
		} else {
			i++
		}
	}

	result = append(result, s[start:])
	return result
}
