package piscine

func IterativeFactorial(nb int) int {
	n := 1
	if nb == 0 {
		return 1
	} else if nb > 0 && nb < 22 {
		for i := 1; i <= nb; i++ {
			n = n * i
		}
	} else {
		return 0
	}
	return n
}
