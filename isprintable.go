package piscine

func IsPrintable(s string) bool {
	for _, c := range s {
		if c < 32 || c > 127 {
			return false
		}
	}
	return true
}
