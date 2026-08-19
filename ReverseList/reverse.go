package reverselist

import "example.com/DSA/cycliclist"

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
