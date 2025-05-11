package smallestsubtreewithallthedeepestnodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	_, node := maxNodeDepth(root)

	return node
}

func maxNodeDepth(root *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}

	lengthLeft, rootLeft := maxNodeDepth(root.Left)
	lengthRight, rootRight := maxNodeDepth(root.Right)

	if lengthLeft > lengthRight {
		return 1 + lengthLeft, rootLeft
	}

	return 1 + lengthRight, rootRight
}
