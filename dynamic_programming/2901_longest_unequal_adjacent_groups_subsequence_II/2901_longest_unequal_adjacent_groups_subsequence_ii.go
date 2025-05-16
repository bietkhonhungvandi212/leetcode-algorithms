package longestunequaladjacentgroupssubsequenceii

func getWordsInLongestSubsequence(words []string, groups []int) []string {

	if len(words) < 1 || len(groups) < 1 || len(words) != len(groups) {
		return []string{}
	}

	n := len(words)
	dp := make([]int, n)
	result := make([]string, 0)
	parent := make(map[int]int)
	max := 1
	dp[0] = 1

	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if groups[i] != groups[j] && isValidHammingDistance(words[i], words[j]) && dp[i] < dp[j]+1 {
				dp[i] = 1 + dp[j]
				parent[i] = j
			}
		}

		if max < dp[i] {
			max = dp[i]
		}
	}

	for i := range n {
		if dp[i] == max {
			ok := true
			for ok == true {
				result = append(result, words[i])
				i, ok = parent[i]
			}

			break
		}
	}

	return reverseStringArray(result)

}

func isValidHammingDistance(wordA string, wordB string) bool {
	if len(wordA) != len(wordB) {
		return false
	}

	count := 0

	for i := range wordA {
		if wordA[i] != wordB[i] {
			count++
		}

		if count == 2 {
			return false
		}
	}

	return count == 1
}

func reverseStringArray(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}
