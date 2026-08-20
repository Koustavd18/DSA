package oddevenlist

import cycliclist "example.com/DSA/CyclicList"

func OddEvenList(head *cycliclist.Node) *cycliclist.Node {
	temp := head
	odd := temp

	if odd.Next == nil {
		return head
	}

	even := odd.Next
	evenList := odd.Next

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next

		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenList

	return head
}
