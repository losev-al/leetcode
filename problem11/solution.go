package problem11

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := make([][]int, 1)
	pointer := &result
	finded := 0
	build(candidates, target, pointer, &finded)
	return result[:len(result)-1]
}

func build(candidates []int, target int, pointer *[][]int, finded *int) {
	for i, candidate := range candidates {
		if 2 * candidate <= target {
			(*pointer)[*finded] = append((*pointer)[*finded], candidate)
			build(candidates[i:], target-candidate, pointer, finded)
		} else if candidate == target {
			*pointer = append(*pointer, make([]int, len((*pointer)[*finded])))
			for j, _ := range (*pointer)[*finded] {
				(*pointer)[*finded + 1][j] = (*pointer)[*finded][j]
			}
			(*pointer)[*finded] = append((*pointer)[*finded], candidate)
			*finded++
		}
	}
	newLen := len((*pointer)[*finded])-1
	if newLen >= 0 {
		(*pointer)[*finded] = (*pointer)[*finded][:newLen]
	}
}
