package piscine

func Unmatch(a []int) int {
	count := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[j] == a[i] {
				count++
			}
		}
		if count%2 != 0 {
			return a[i]
		}
		count = 0
	}
	return -1
}
