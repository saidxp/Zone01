package piscine

func StrLen(s string) int {
	count := 0
	str := []rune(s)
	for i := 0; i < len(str); i++ {
		count++
	}
	return count
}
