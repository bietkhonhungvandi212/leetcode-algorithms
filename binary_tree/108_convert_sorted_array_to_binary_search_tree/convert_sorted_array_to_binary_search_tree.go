package convertsortedarraytobinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	return BuildBST(nums, 0, n-1)
}

func BuildBST(nums []int, start int, end int) *TreeNode {
	mid := start + (end-start)/2
	root := &TreeNode{Val: nums[mid], Left: &TreeNode{}, Right: &TreeNode{}}

	if start < mid {
		root.Left = BuildBST(nums, start, mid-1)
	}

	if mid < end {
		root.Right = BuildBST(nums, mid+1, end)
	}

	return root
}
