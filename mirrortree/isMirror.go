package mirrortree

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func FindMirror(t *TreeNode) bool {

	if t.Left == nil && t.Right == nil {
		return true
	}

	return isMirror(t.Left, t.Right)

}

func isMirror(t1 *TreeNode, t2 *TreeNode) bool {

	if t1 == nil && t2 == nil {
		return true
	}

	if t1 == nil || t2 == nil {
		return false
	}

	return t1.Value == t2.Value && isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
}
