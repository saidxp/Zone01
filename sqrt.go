package piscine

func Sqrt(nb int) int {
	count := 0
	if nb <= 0 {
		return 0
	}
	for i := 1; i*i <= nb; i++ {
		count++
		if count*count == nb {
			return count
		}
	}
	return 0
}
