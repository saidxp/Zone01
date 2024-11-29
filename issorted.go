package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	inc := true
	dec := true
	for i := 0; i+1 < len(a); i++ {
		if f(a[i], a[i+1]) > 0 {
			inc = false
		} else if f(a[i], a[i+1]) < 0 {
			dec = false
		}
	}
	return inc || dec
}

func F(a, b int) int {
	min := a - b
	if min > 0 {
		return 1
	} else if min < 0 {
		return -1
	}
	return 0
}
