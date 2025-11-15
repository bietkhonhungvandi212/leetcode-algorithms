package floodfill

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	rows := len(image)
	cols := len(image[0])
	target := image[sr][sc]

	if target == color {
		return image
	}

	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{sr, sc})
	image[sr][sc] = color

	for len(queue) > 0 {
		r, c := queue[0][0], queue[0][1]
		queue = queue[1:]

		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && image[nr][nc] == target {
				image[nr][nc] = color
				queue = append(queue, [2]int{nr, nc})
			}
		}
	}

	return image
}
