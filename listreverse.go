package piscine

func ListReverse(l *List) {
	var temp *NodeL
	l.Tail = l.Head
	current := l.Head
	for current != nil {
		nex := current.Next
		current.Next = temp
		temp = current
		current = nex
	}
	l.Head = temp
}
