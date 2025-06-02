package snakesandladders

import (
	"fmt"
)

func snakesAndLadders(board [][]int) int {
	arrays := buildArray(board)
	length := len(board)
	target := length * length
	fmt.Println(arrays)

	queue := []int{1}
	visit := make(map[int]bool)
	visit[1] = true
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		for range size {
			current := queue[0]
			queue = queue[1:]

			if current == target {
				return steps
			}

			// Roll dice: 1 to 6
			for dice := 1; dice <= 6; dice++ {
				next := current + dice
				if next > target {
					continue
				}

				// Check for snake or ladder
				if arrays[next-1] != -1 {
					next = arrays[next-1]
				}

				if !visit[next] {
					visit[next] = true
					queue = append(queue, next)
				}
			}
		}
		steps++
	}

	return -1 // Target unreachable
}

func buildArray(board [][]int) []int {
	length := len(board)
	result := make([]int, length*length)
	index := 0
	for i := length - 1; i >= 0; i-- {
		for range length {
			if (length-1-i)%2 == 0 {
				result[index] = board[i][index%length]
			} else {
				result[index] = board[i][length-1-(index%length)]
			}
			index++
		}
	}

	return result
}
