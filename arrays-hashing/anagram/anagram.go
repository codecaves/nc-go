package anagram

import "sort"

func sortRunesToString(s string) string {
	runes := []rune(s)

	sort.Slice(runes, func(i int, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}