package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	words := make(map[string][]string)
	for _, str := range strs {
		key := SortString(str)
		words[key] = append(words[key], str)
	}
	result := make([][]string, 0, len(words))
	for _, strings := range words {
		result = append(result, strings)
	}
	return result
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
