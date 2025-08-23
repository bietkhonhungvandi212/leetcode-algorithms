package longestcommonprefix

import (
	"bytes"
	"sort"
)

func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)
	first := strs[0]
	last := strs[len(strs)-1]

	var buffer bytes.Buffer
	for i := 0; i < min(len(first), len(last)); i++ {
		if first[i] != last[i] {
			break
		}

		buffer.WriteByte(first[i])
	}

	return buffer.String()
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
