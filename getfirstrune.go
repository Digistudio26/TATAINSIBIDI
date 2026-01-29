package main




func GetFirstRune(s string) rune {
	for _, r := range s {
		return r
	}
	return 0
}

func main() {
	fmt.Println(GetFirstRune("kood"))
}
