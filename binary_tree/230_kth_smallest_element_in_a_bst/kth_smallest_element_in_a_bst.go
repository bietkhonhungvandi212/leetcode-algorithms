package kthsmallestelementinabst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Track struct {
	Count int
	Res   int
}

func kthSmallest(root *TreeNode, k int) int {
	var result []int

	traverseInOrderTree(root, &result)

	return result[k-1]
}

func traverseInOrderTree(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	traverseInOrderTree(root.Left, result)
	*result = append(*result, root.Val)
	traverseInOrderTree(root.Right, result)
}

func kthSmallestV2(root *TreeNode, k int) int {
	track := Track{
		Count: 0,
		Res:   0,
	}

	countInOrderTree(root, k, &track)

	return track.Res
}

func countInOrderTree(root *TreeNode, k int, track *Track) {
	if root == nil || track.Res != 0 {
		return
	}

	countInOrderTree(root.Left, k, track)
	*&track.Count++
	if *&track.Count == k {
		*&track.Res = root.Val
	}

	countInOrderTree(root.Right, k, track)
}
