package main

func lengthOfLongestSubstring(s string) int {
	left := 0
	right := 0
	max := 0
	set := StringSet{}
	length := len(s)

	for right < length {
		for set.has(s[right]) {
			set.remove(s[left])
			left++
		}

		set.add(s[right])
		max = MaxInt(max, len(set))

		right++
	}

	return max
}

type StringSet map[byte]bool

// Adds an character to the set
func (s StringSet) add(character byte) {
	s[character] = true
}

// Removes an character from the set
func (s StringSet) remove(character byte) {
	delete(s, character)
}

func (s StringSet) has(character byte) bool {
	_, ok := s[character]
	return ok
}

func MaxInt(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
