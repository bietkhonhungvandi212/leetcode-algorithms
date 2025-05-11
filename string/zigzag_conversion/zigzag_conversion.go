package main

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	arr2D := make([][]byte, numRows)
	for i := range arr2D {
		arr2D[i] = make([]byte, 0)
	}

	row := 0
	d := 1
	for i := range s {
		arr2D[row] = append(arr2D[row], s[i])
		if row == 0 {
			d = 1
		}
		if row == numRows-1 {
			d = -1
		}

		row += d
	}

	var result strings.Builder
	for _, bytes := range arr2D {
		result.Write(bytes)
	}
	return result.String()

}
