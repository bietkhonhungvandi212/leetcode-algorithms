package maximizethenumberoftargetnodesafterconnectingtreesi

import "fmt"

func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {
	//1 Build Graph for 2 trees
	graphEdges1 := buildGraph(edges1)
	graphEdges2 := buildGraph(edges2)
	result := make([]int, len(edges1)+1)

	//2 Looping tree 1, with edges which is less than or equal k paths is to count the targets of each nodes
	fmt.Println(graphEdges1)
	fmt.Println(graphEdges2)
	for i := range len(edges1) + 1 {
		result[i] = bfs(graphEdges1, i, k)
	}
	fmt.Println(result)

	//3 Looping tree 2, to find the the node has the most targets
	maxPaths := 0
	for i := range len(edges2) + 1 {
		maxPaths = max(maxPaths, bfs(graphEdges2, i, k-1))
	}

	//4 Result is the sum of targets of each element at first tree with the max of second tree
	for i := range result {
		result[i] += maxPaths
	}
	fmt.Println(result)

	return result
}

func max(num1, num2 int) int {
	if num1 < num2 {
		return num2
	}

	return num1
}

func buildGraph(edges [][]int) map[int][]int {
	graph := make(map[int][]int, len(edges))
	for i := range edges {
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
	}

	return graph
}

func bfs(graph map[int][]int, node, k int) int {
	if k < 0 {
		return 0
	}

	queue := []int{node}
	visit := make([]bool, len(graph))
	visit[node] = true
	cnt := 1 // Count the starting node at distance 0

	for level := 0; level < k && len(queue) > 0; level++ {
		size := len(queue)
		for i := 0; i < size; i++ {
			currentNode := queue[0] // Take from FRONT (queue behavior)
			queue = queue[1:]       // Remove from FRONT

			for _, neighbor := range graph[currentNode] {
				if !visit[neighbor] {
					visit[neighbor] = true
					queue = append(queue, neighbor)
					cnt++
				}
			}
		}
	}

	return cnt
}
