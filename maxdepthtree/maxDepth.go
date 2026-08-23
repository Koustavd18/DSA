package maxdepthtree

import (
	"math"

	"example.com/DSA/mirrortree"
)

func MaxDepth(t *mirrortree.TreeNode) int {

	if t == nil {
		return 0
	}

	left := 1 + MaxDepth(t.Left)
	right := 1 + MaxDepth(t.Right)

	ans := math.Max(float64(left), float64(right))

	return int(ans)

}
