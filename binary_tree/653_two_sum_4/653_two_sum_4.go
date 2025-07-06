package twosum4

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	stack := make([]*TreeNode, 0)
	result := make([]int, 0)

	curr := root

	for curr != nil || len(stack) > 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, top.Val)
			if top.Right != nil {
				curr = top.Right
			} else {
				curr = nil
			}
		}

	}

	left := 0
	right := len(result) - 1

	for left < right {
		if result[left]+result[right] == k {
			return true
		}

		if result[left]+result[right] > k {
			right--
		}

		if result[left]+result[right] < k {
			left++
		}

	}

	return false
}

func findTargetV2(root *TreeNode, k int) bool {
	mapper := make(map[int]struct{})

	var dfs func(root *TreeNode, k int) bool
	dfs = func(root *TreeNode, k int) bool {
		if root == nil {
			return false
		}

		subtract := k - root.Val

		if _, ok := mapper[root.Val]; ok {
			return true
		}
		mapper[subtract] = struct{}{}

		return dfs(root.Left, k) || dfs(root.Right, k)
	}

	return dfs(root, k)
}
