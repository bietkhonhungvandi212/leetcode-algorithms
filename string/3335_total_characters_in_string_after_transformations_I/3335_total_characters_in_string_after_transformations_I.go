package totalcharactersinstringaftertransformationsi

func lengthAfterTransformations(s string, t int) int {
	mod := 1000000007
	freq := make([]int, 26)
	count := 0

	for i := range s {
		index := (s[i] - 'a') % 26
		freq[index]++
	}

	for _ = t; t > 0; t-- {
		freqZ := freq[25]
		for i := 24; i >= 0; i-- {
			freq[i+1] = freq[i]
		}

		freq[0] = freqZ
		freq[1] = (freqZ + freq[1]) % mod
	}

	for i := range freq {
		if freq[i] > 0 {
			count = (count + freq[i]) % mod
		}
	}

	return count
}
