package addlist

import cycliclist "example.com/DSA/CyclicList"

func Add(l1, l2 *cycliclist.Node) *cycliclist.Node {

	ans := &cycliclist.Node{}
	pointer := ans
	var carry, sum = 0, 0

	_l1 := l1
	_l2 := l2

	for _l1 != nil || _l2 != nil {
		sum = carry

		if _l1 != nil {
			sum += _l1.Value
			_l1 = _l1.Next
		}
		if _l2 != nil {
			sum += _l2.Value
			_l2 = _l2.Next
		}
		carry = int(sum / 10)
		pointer.Next = &cycliclist.Node{Value: sum % 10}
		pointer = pointer.Next

	}

	if carry == 1 {
		pointer.Next = &cycliclist.Node{Value: carry}
	}

	return ans.Next
}
