package main

func longestSubstring(s string, k int) int {
	return devideSubsString(s, k)
}

func devideSubsString(s string, k int) int {
	if len(s) < k {
		return 0
	}

	frequency := make(map[byte]int)

	for i := range s {
		frequency[s[i]]++
	}

	for i := range s {
		if frequency[s[i]] < k {
			left := devideSubsString(s[:i], k)
			right := devideSubsString(s[i+1:], k)

			return MaxInt(left, right)
		}
	}

	return len(s)

}

func MaxInt(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
