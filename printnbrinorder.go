package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	nbr := Itoi(n)
	str := []rune(nbr)
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(str)-1-i; j++ {
			if str[j] > str[j+1] {
				swapor := str[j]
				str[j] = str[j+1]
				str[j+1] = swapor
			}
		}
	}
	for _, c := range str {
		z01.PrintRune(c)
	}
}

func Itoi(nbr int) string {
	if nbr == 0 {
		return "0"
	}
	result := ""
	for nbr > 0 {
		result = string(nbr%10+48) + result
		nbr /= 10
	}
	return result
}
