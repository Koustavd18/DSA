package reverselist

import cycliclist "example.com/DSA/CyclicList"

func Reverse(head *cycliclist.Node) *cycliclist.Node {

	var prev *cycliclist.Node = nil
	cur := head

	for cur != nil {
		next := cur.Next

		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
