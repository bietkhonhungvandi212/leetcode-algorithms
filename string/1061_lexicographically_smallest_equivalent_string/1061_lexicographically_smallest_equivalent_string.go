package lexicographicallysmallestequivalentstring

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	characters := make([]int, 26)
	result := make([]byte, len(baseStr))

	var find func(a int) int
	find = func(a int) int {
		if characters[a] != a {
			characters[a] = find(characters[a]) // Path compression
		}
		return characters[a]
	}

	for i := range characters {
		characters[i] = i
	}

	for i := range s1 {
		index1 := int(s1[i] - 'a')
		index2 := int(s2[i] - 'a')

		root1 := find(index1)
		root2 := find(index2)

		if root1 == root2 {
			continue
		}

		if root1 < root2 {
			characters[root2] = root1
		} else {
			characters[root1] = root2
		}
	}

	for i := range baseStr {
		index := int(baseStr[i] - 'a')
		smallestRoot := find(index)
		result[i] = byte(smallestRoot) + 'a'
	}

	return string(result)
}
