package piscine

func Rot14(s string) string {
	str := []rune(s)
	for i := 0; i < len(str); i++ {
		if str[i] >= 'a' && str[i] <= 'z' {
			str[i] = 'a' + (str[i]-'a'+14)%26
		} else if str[i] >= 'A' && str[i] <= 'Z' {
			str[i] = 'A' + (str[i]-'A'+14)%26
		}
	}
	return string(str)
}
