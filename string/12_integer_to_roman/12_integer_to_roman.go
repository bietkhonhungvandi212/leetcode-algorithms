package integertoroman

func intToRoman(num int) string {
	numbers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	stack := make([]byte, 0)

	for i := range numbers {
		for num-numbers[i] > 0 {
			for j := range romans[i] {
				stack = append(stack, romans[i][j])
			}
			num -= numbers[i]
		}
	}

	return string(stack)
}
