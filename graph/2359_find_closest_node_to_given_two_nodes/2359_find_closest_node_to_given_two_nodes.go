package findclosestnodetogiventwonodes

import (
	"fmt"
	"math"
)

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	//1 initial distance from node1 -> all nodes, node2 -> all nodes
	distance1 := bfs(edges, node1)
	distance2 := bfs(edges, node2)

	fmt.Println(distance1)
	fmt.Println(distance2)

	var findMinimumOfMaxBetweenDistances func(min *int, index *int, currentIndex int, distance1, distance2 int)
	findMinimumOfMaxBetweenDistances = func(min *int, index *int, currentIndex int, distance1, distance2 int) {
		max := max(distance1, distance2)
		if max < *min {
			*min = max
			*index = currentIndex
		}
	}

	//2 Find the minimum distance from node1 -> all nodes
	min := math.MaxInt
	indexMin := -1
	for i := range len(edges) {
		// Skip nodes that are unreachable from either node1 or node2
		if distance1[i] == -1 || distance2[i] == -1 {
			continue
		}
		findMinimumOfMaxBetweenDistances(&min, &indexMin, i, distance1[i], distance2[i])
	}

	return indexMin
}

func bfs(edges []int, node int) []int {
	result := make([]int, len(edges))
	// Initialize all distances to -1 (unreachable)
	for i := range result {
		result[i] = -1
	}

	if node < 0 || node >= len(edges) {
		return result
	}

	visit := make([]bool, len(edges))

	currentNode := node
	count := 0

	for currentNode != -1 && !visit[currentNode] {
		result[currentNode] = count
		visit[currentNode] = true
		count++
		currentNode = edges[currentNode]
	}

	return result
}

func max(num1, num2 int) int {
	if num1 < num2 {
		return num2
	}

	return num1
}
