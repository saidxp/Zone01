package piscine

func ListClear(l *List) {
	current := l.Head
	for current != nil {
		cur := current.Next
		current.Next = nil
		current = cur
	}
	l.Head = nil
	l.Tail = nil
}
