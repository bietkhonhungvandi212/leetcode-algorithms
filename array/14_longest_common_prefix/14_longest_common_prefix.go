package longestcommonprefix

import (
	"math"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	min := math.MaxInt
	if len(strs) == 1 {
		return strs[0]
	}
	for i := range strs {
		if len(strs[i]) == 0 {
			return ""
		}

		min = Min(min, len(strs[i]))
	}

	var result strings.Builder
	var tmp byte
	for i := range min {
		tmp = strs[0][i]

		for j := range strs {
			if strs[j][i] != tmp {
				return result.String()
			}
		}
		result.WriteByte(strs[0][i])
	}

	return result.String()
}

func Min(a, b int) int {
	if a > b {
		return a
	}

	return b
}
