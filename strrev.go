package piscine

func StrRev(s string) string {
	str := []rune(s)
	lkhit := ""

	for i := len(str) - 1; i >= 0; i-- {
		lkhit = lkhit + string(str[i])
	}
	return lkhit
}
