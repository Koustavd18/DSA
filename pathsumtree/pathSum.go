package pathsumtree

import "example.com/DSA/mirrortree"

func PathSum(root *mirrortree.TreeNode, t int) bool {

	if root == nil {
		return false
	}
	return hasSum(root, t, 0)
}

func hasSum(root *mirrortree.TreeNode, t int, cur int) bool {

	cur += root.Value

	if cur == t && root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left != nil {
		if hasSum(root.Left, t, cur) {
			return true
		}
	}

	if root.Right != nil {
		if hasSum(root.Right, t, cur) {
			return true
		}
	}

	return false
}
