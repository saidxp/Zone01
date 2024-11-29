package piscine

func LastRune(s string) rune {
	str := []rune(s)
	r := str[len(str)-1]
	return r
}
