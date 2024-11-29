package piscine

func UltimateDivMod(a *int, b *int) {
	mod := *a / *b
	div := *a % *b
	*a = mod
	*b = div
}
