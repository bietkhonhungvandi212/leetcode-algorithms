package excelsheetcolumnnumber

func titleToNumber(columnTitle string) int {
	result := 0
	for _, val := range columnTitle {
		d := val - 'A' + 1

		result = result*26 + int(d)
	}
	return result
}
