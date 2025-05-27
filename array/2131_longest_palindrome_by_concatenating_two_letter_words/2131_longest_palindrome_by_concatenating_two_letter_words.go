package longestpalindromebyconcatenatingtwoletterwords

import "fmt"

func longestPalindrome(words []string) int {
	frequencies := make(map[string]int)
	for i := range words {
		frequencies[words[i]]++
	}

	fmt.Println(frequencies)

	midPalidrome := 0
	count := 0

	for key := range frequencies {
		frequency := frequencies[key]
		s1 := key[0]
		s2 := key[1]

		if s1 == s2 {
			count += (frequency / 2) * 4
			if frequency%2 == 1 {
				midPalidrome = 1
			}
		} else if s1 < s2 {
			reversedWordStr := string(s2) + string(s1)
			if reverseFrequency, ok := frequencies[reversedWordStr]; ok {
				count += min(frequency, reverseFrequency) * 4
			}
		}
	}

	return count + midPalidrome*2
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}

	return num2
}
