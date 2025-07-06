package minimumabsolutedifferenceinbst

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	min := math.MaxInt
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

	fmt.Println(result)

	for i := range len(result) - 1 {
		subtract := result[i+1] - result[i]
		if min > subtract && subtract > 0 {
			min = subtract
		}
	}

	return min
}
