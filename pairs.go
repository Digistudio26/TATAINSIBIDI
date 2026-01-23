package main

func Pairs() string {
	result := ""

	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			if result != "" {
				result += ", "
			}
			result += string('0'+i/10) + string('0'+i%10) +
				" " +
				string('0'+j/10) + string('0'+j%10)
		}
	}

	return result
}