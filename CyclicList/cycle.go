package cycliclist

type Node struct {
	Value int
	Next  *Node
}

func IsCycle(head *Node) bool {

	p1 := head
	p2 := head

	for p1 != nil && p1.Next != nil {

		p2 = p2.Next
		p1 = p1.Next.Next

		if p1 == p2 {
			return true
		}

	}

	return false
}
