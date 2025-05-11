package main

import "reflect"

func checkInclusion(s1 string, s2 string) bool {
	length1 := len(s1)
	length2 := len(s2)

	if length2 < length1 {
		return false
	}

	appearanceString1 := StringSet{}
	appearanceString2 := StringSet{}

	for i := range s1 {
		appearanceString1[s1[i]]++
		appearanceString2[s2[i]]++
	}

	if reflect.DeepEqual(appearanceString1, appearanceString2) {
		return true
	}

	left := 0
	for right := length1; right < length2; right++ {
		appearanceString2[s2[right]]++
		appearanceString2[s2[left]]--

		if appearanceString2[s2[left]] == 0 {
			appearanceString2.remove(s2[left])
		}

		left++

		if reflect.DeepEqual(appearanceString1, appearanceString2) {
			return true
		}
	}

	return false
}

type StringSet map[byte]int

func (s StringSet) remove(character byte) {
	delete(s, character)
}
