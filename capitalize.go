package piscine

func Capitalize(s string) string {
	str := []rune(s)
	for i, char := range str {
		if char >= 'A' && char <= 'Z' {
			str[i] += 32
		}
	}
	if str[0] >= 'a' && str[0] <= 'z' {
			str[0] -= 32
	}
	for i, r := range str {
		if i+1 < len(str) && (r >= 32 && r <= 47) && (str[i+1] >= 'a' && str[i+1] <= 'z') {
			str[i+1] -= 32
		}
	}
	return string(str)
}
