package piscine

func IterativePower(nb int, power int) int {
	var n int
	n = 1

	if power == 0 {
		return 1
	} else if nb == 0 {
		return 0
	}
	if power < 0 {
		return 0
	}
	for i := 1; i <= power; i++ {
		n = n * nb
	}
	return n
}
