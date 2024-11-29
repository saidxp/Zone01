package piscine

func ListSize(l *List) int {
	count := 0
	if l.Head == nil {
		return 0
	}
	current := l.Head
	for current != nil {
		current = current.Next
		count++
	}
	return count
}
