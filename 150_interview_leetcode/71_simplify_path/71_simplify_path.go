package simplifypath

import "strings"

func simplifyPath(path string) string {
	tokens := strings.Split(path, "/")
	stack := make([]string, 0)

	for _, val := range tokens {
		if val == "" || val == "." {
			continue
		}

		if val == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, val)
		}
	}

	result := "/" + strings.Join(stack, "/")
	if len(stack) == 0 {
		return "/"
	}
	return result
}
