package piscine

func Swap(a *int, b *int) {
	var swap int
	swap = *a
	*a = *b
	*b = swap
}
