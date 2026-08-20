package removenth

import cycliclist "example.com/DSA/CyclicList"

func LengthTraverse(l1 *cycliclist.Node, n int) *cycliclist.Node {

	temp := l1
	len := 0

	for temp != nil {
		len += 1
		temp = temp.Next
	}

	if len == n {
		return l1.Next
	}

	len -= n + 1
	first := l1

	for len > 0 {
		first = first.Next
		len -= 1
	}

	first.Next = first.Next.Next

	return l1

}

func RemoveNode(head *cycliclist.Node, n int) *cycliclist.Node {

	temp := &cycliclist.Node{Value: 0}
	temp.Next = head

	slow := temp
	fast := temp.Next

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return temp.Next
}
