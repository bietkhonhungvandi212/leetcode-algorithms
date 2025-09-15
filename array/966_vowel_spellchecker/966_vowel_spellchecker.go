package vowelspellchecker

import (
	"fmt"
	"strings"
)

func spellchecker(wordlist []string, queries []string) []string {
	insenMatch := make(map[string]string)
	vowels := make(map[string]string)
	match := make(map[string]bool)

	for _, word := range wordlist {
		match[word] = true
		lower := strings.ToLower(word)
		if _, ok := insenMatch[lower]; !ok {
			insenMatch[lower] = word
		}
		vowelKey := makeVowel(lower)
		fmt.Println(vowels)
		if _, ok := vowels[vowelKey]; !ok {
			vowels[vowelKey] = word
		}
	}

	result := make([]string, len(queries))
	for i, query := range queries {
		if match[query] {
			result[i] = query
			continue
		}
		lower := strings.ToLower(query)
		if word, ok := insenMatch[lower]; ok {
			result[i] = word
			continue
		}

		vowelKey := makeVowel(lower)
		if word, ok := vowels[vowelKey]; ok {
			result[i] = word
			continue
		}

		result[i] = ""
	}

	return result
}

func makeVowel(s string) string {
	bytes := []byte(s)
	for i, b := range bytes {
		if isVowel(b) {
			bytes[i] = '*'
		}
	}
	return string(bytes)
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}
