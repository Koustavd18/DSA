package mergesortedlist

import "container/list"

func MergeLists(l1, l2 *list.List) *list.List {

	result := list.New()

	e1 := l1.Front()
	e2 := l2.Front()

	for e1 != nil && e2 != nil {

		v1 := e1.Value.(int)
		v2 := e2.Value.(int)

		if v1 <= v2 {
			result.PushBack(v1)
			e1.Next()
		} else {
			result.PushBack(v2)
			e2.Next()
		}

	}

	for e1 != nil {
		result.PushBack(e1.Value.(int))
		e1.Next()
	}

	for e2 != nil {
		result.PushBack(e2.Value.(int))
		e2.Next()
	}

	return result
}
