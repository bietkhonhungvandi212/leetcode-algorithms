package main

func characterReplacement(s string, k int) int {
	left := 0
	length := len(s)
	maxCount := 0
	res := 0
	countMap := StringSet{}

	for right := range length {
		countMap[s[right]]++

		maxCount = MaxInt(maxCount, countMap[s[right]])

		for right-left+1-maxCount > k {
			countMap[s[left]]--
			left++

			for k := range countMap {
				maxCount = MaxInt(countMap[k], maxCount)
			}
		}

		res = MaxInt(res, right-left+1)
	}

	return res
}

func MaxInt(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

type StringSet map[byte]int

func (s StringSet) has(character byte) bool {
	_, ok := s[character]
	return ok
}
