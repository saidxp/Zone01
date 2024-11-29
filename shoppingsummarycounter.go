package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	m := make(map[string]int)

	if len(str) == 0 {

		m[""] = 1
		return m
	}

	onlySpaces := true
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			onlySpaces = false
			break
		}
	}

	if onlySpaces {

		m[""] = len(str) + 1
		return m
	}

	words := []string{}
	from := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {

			if i > from {
				words = append(words, str[from:i])
			} else {
				words = append(words, "")
			}
			from = i + 1
		}
	}

	if len(str) > from {
		words = append(words, str[from:])
	} else {
		words = append(words, "")
	}

	for _, word := range words {
		m[word]++
	}

	return m
}
