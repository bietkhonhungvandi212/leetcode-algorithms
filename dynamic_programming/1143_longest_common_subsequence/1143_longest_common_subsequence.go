package longestcommonsubsequence

func longestCommonSubsequence(text1 string, text2 string) int {
	arr2D := make([][]int, len(text1)+1)
	for i := range arr2D {
		arr2D[i] = make([]int, len(text2)+1)
	}

	for i := range text1 {
		for j := range text2 {
			if text1[i] == text2[j] {
				arr2D[i+1][j+1] = 1 + arr2D[i][j]
			} else {
				arr2D[i+1][j+1] = max(arr2D[i+1][j], arr2D[i][j+1])
			}
		}
	}

	return arr2D[len(text1)][len(text2)]

}

func longestCommonSubsequence_Topdown(text1 string, text2 string) int {
	arr2D := make([][]int, len(text1))
	for i := range arr2D {
		arr2D[i] = make([]int, len(text2))
	}

	var countCommonSubsequence func(i, j int) int
	countCommonSubsequence = func(i, j int) int {
		if i >= len(text1) || j >= len(text2) {
			return 0
		}

		if arr2D[i][j] != 0 {
			return arr2D[i][j]
		}

		if text1[i] == text2[j] {
			arr2D[i][j] = 1 + countCommonSubsequence(i+1, j+1)
			return arr2D[i][j]
		}

		arr2D[i][j] = max(countCommonSubsequence(i+1, j), countCommonSubsequence(i, j+1))
		return arr2D[i][j]
	}

	return countCommonSubsequence(0, 0)
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
