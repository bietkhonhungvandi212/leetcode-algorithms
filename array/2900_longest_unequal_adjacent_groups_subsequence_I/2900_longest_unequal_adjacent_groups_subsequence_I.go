package longestunequaladjacentgroupssubsequencei

func getLongestSubsequence(words []string, groups []int) []string {
	if len(words) == 1 || len(groups) == 1 || len(words) != len(groups) {
		return []string{words[0]}
	}

	result := make([]string, 0)
	n := len(words)

	for i := n - 2; i >= 0; i-- {
		if groups[i] != groups[i+1] {
			if len(result) == 0 {
				result = append(result, words[i+1])
				result = append(result, words[i])
			} else {
				result = append(result, words[i])
			}

		}

	}

	if len(result) == 0 {
		return []string{words[0]}
	}

	return reverseStringArray(result)
}

func reverseStringArray(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}
